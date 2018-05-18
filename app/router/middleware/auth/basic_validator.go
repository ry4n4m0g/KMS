package auth

import (
	API "github.com/epointpayment/key_management_system/app/services/api"
	"github.com/labstack/echo"
)

// BasicValidator is a validator used for basic auth middleware
func BasicValidator(username, password string, c echo.Context) (isValid bool, err error) {
	sa := API.New()

	isValid, err = sa.DoAuth(username, password)
	return
}
