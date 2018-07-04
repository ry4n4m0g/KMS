package controllers

import (
	"net/http"
	"strconv"

	Key "github.com/epointpayment/key_management_system/app/services/key"
	KeyGen "github.com/epointpayment/key_management_system/app/services/keygen"

	"github.com/labstack/echo"
)

// GetKey retrieves a key from the database
func (co Controllers) GetKey(c echo.Context) error {
	// Get key ID
	keyID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		err = Key.ErrKeyInvalid
		return err
	}

	// Initialize key service
	userID := c.Get("userID").(int)
	ks := Key.New(userID)

	// Get key by key ID
	res, err := ks.Get(keyID)
	if err != nil {
		return err
	}

	// Send results
	return SendResponse(c, http.StatusOK, res)
}

// GenerateKey creates a new key and stores it in the database
func (co Controllers) GenerateKey(c echo.Context) error {
	// Initialize key service
	userID := c.Get("userID").(int)
	ks := Key.New(userID)

	// Get key algorithm
	keyAlgorithm := c.QueryParam("algorithm")

	// Get key length
	queryKeyLength := c.QueryParam("length")
	keyLength, err := strconv.Atoi(queryKeyLength)
	if queryKeyLength != "" && err != nil {
		err = KeyGen.ErrKeyLengthInvalid
		return err
	}

	// Generate key
	res, err := ks.Generate(keyAlgorithm, keyLength)
	if err != nil {
		return err
	}

	// Send results
	return SendResponse(c, http.StatusOK, res)
}
