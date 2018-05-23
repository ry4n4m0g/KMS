package key

import "errors"

var (
	// ErrKeyInvalid is an error given when the key is not valid
	ErrKeyInvalid = errors.New("Invalid Key ID")

	// ErrKeyNotFound is an error for a non-existent key
	ErrKeyNotFound = errors.New("Key ID not found")
)
