package keygen

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/jmcvetta/randutil"
)

// KeyGenService is a service generates a random hash
type KeyGenService struct{}

// New creates an instance of the service
func New() *KeyGenService {
	return &KeyGenService{}
}

// Generate creates a random hash
func (ks *KeyGenService) Generate() (hash string, err error) {
	strRandomText, err := ks.generateRandomString()
	if err != nil {
		return
	}

	hash = ks.hasher(strRandomText)

	return
}

// hasher hashes a string
func (ks *KeyGenService) hasher(str string) string {
	h := sha256.New()

	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
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
