package keygen

import "errors"

var (
	// ErrKeyAlgorithmInvalid is an error given when the key algorithm is not in the allowed list
	ErrKeyAlgorithmInvalid = errors.New("Invalid Key Algorithm")

	// ErrKeyLengthInvalid is an error given when the key length is invalid or out of the specified range
	ErrKeyLengthInvalid = errors.New("Invalid Key Length")
)
