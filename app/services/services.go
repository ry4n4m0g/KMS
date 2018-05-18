package services

import (
	"github.com/epointpayment/key_management_system/app/database"

	dbx "github.com/go-ozzo/ozzo-dbx"
)

var db *dbx.DB

type Services struct{}

func New(DB *database.Database) error {
	db = DB.GetInstance()

	return nil
}
