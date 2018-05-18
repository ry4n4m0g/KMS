package models

import "time"

// Keys is an array of Key entries
type Keys []Key

// Key contains information about a generated a key
type Key struct {
	ID          int       `json:"id"`
	Key         string    `json:"key"`
	DateCreated time.Time `json:"-"`
}

// TableName gets the name of the database table
func (k Key) TableName() string {
	return "keys"
}
