package key

import "errors"

var (
	// ErrKeyNotFound is an error for a non-existent key
	ErrKeyNotFound = errors.New("Key not found")
)
