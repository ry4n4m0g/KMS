package keygen

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"

	"github.com/go-ozzo/ozzo-validation"
	"github.com/jmcvetta/randutil"
)

const (
	SHA256 = "SHA256"
)

// KeyGenService is a service generates a random hash
type KeyGenService struct {
	Algorithm string
	Length    int
}

// Validate checks the configuration for invalid values
func (ks KeyGenService) Validate() (err error) {
	// Check algorithm
	err = validation.Validate(ks.Algorithm,
		validation.In(SHA256),
	)
	if err != nil {
		err = ErrKeyAlgorithmInvalid
		return
	}

	// Check length
	err = validation.Validate(ks.Length,
		validation.Min(8),
		validation.Max(64),
	)
	if err != nil {
		err = ErrKeyLengthInvalid
		return
	}

	return
}

// New creates an instance of the service
func New(algorithm string, length int) *KeyGenService {
	return &KeyGenService{
		Algorithm: strings.ToUpper(algorithm),
		Length:    length,
	}
}

// Generate creates a random hash
func (ks *KeyGenService) Generate() (hash string, err error) {
	err = ks.Validate()
	if err != nil {
		return
	}

	strRandomText, err := ks.generateRandomString()
	if err != nil {
		return
	}

	hash = ks.hasher(strRandomText)

	return
}

// hasher hashes a string
func (ks *KeyGenService) hasher(str string) (s string) {
	switch ks.Algorithm {
	case SHA256:
		fallthrough
	default:
		h := sha256.New()
		h.Write([]byte(str))
		s = hex.EncodeToString(h.Sum(nil))
	}

	if ks.Length > 0 {
		s = s[:ks.Length]
	}

	return
}

// generateRandomString creates a random string
func (ks *KeyGenService) generateRandomString() (str string, err error) {
	str, err = randutil.String(256, randutil.Ascii)
	if err != nil {
		return
	}
	str += "." + time.Now().UTC().String()

	return
}
