package user

import "errors"

var (
	// ErrUserNotFound is an error for a non-existent user
	ErrUserNotFound = errors.New("User not found")

	// ErrUserExists is an error shown when trying to create a user that already exists
	ErrUserExists = errors.New("User already exists")
)
