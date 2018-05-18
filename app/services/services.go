package services

import (
	"github.com/epointpayment/key_management_system/app/database"
	"github.com/epointpayment/key_management_system/app/services/key"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

// db is the database handler
var db *dbx.DB

// Services boots application-specific services
type Services struct{}

// New starts the service setup process
func New(DB *database.Database) error {
	// Attach the database handler to service
	db = DB.GetInstance()
	key.DB = db

	return nil
}
