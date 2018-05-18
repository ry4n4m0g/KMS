package controllers

import (
	"net/http"
	"strconv"

	Key "github.com/epointpayment/key_management_system/app/services/key"
	"github.com/labstack/echo"
)

// GetKey retrieves a key from the database
func (co Controllers) GetKey(c echo.Context) error {
	keyID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return err
	}

	ks := Key.New()

	res, err := ks.Get(keyID)
	if err != nil {
		return err
	}

	return SendResponse(c, http.StatusOK, res)
}

// GenerateKey creates a new key and stores it in the database
func (co Controllers) GenerateKey(c echo.Context) error {
	ks := Key.New()

	res, err := ks.Generate()
	if err != nil {
		return err
	}

	return SendResponse(c, http.StatusOK, res)
}
