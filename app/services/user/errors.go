package user

import "errors"

var (
	// ErrUserNotFound is an error for a non-existent user
	ErrUserNotFound = errors.New("User not found")
)
