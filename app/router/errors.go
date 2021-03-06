package router

import (
	"net/http"

	"github.com/epointpayment/key_management_system/app/controllers"
	API "github.com/epointpayment/key_management_system/app/services/api"
	Key "github.com/epointpayment/key_management_system/app/services/key"
	KeyGen "github.com/epointpayment/key_management_system/app/services/keygen"
	User "github.com/epointpayment/key_management_system/app/services/user"

	"github.com/labstack/echo"
)

// appendErrorHandler handles errors for the router
func (r *Router) appendErrorHandler() {
	r.e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		message := err.Error()
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			message = he.Message.(string)
		}

		// Override status code based on error responses
		switch message {
		// API Service
		case API.ErrInvalidCredentials.Error():
			code = http.StatusForbidden

		// Key Service
		case Key.ErrKeyNotFound.Error():
			code = http.StatusNotFound
		case Key.ErrKeyInvalid.Error():
			code = http.StatusBadRequest

			// KeyGen Service
		case KeyGen.ErrKeyAlgorithmInvalid.Error():
			code = http.StatusBadRequest
		case KeyGen.ErrKeyLengthInvalid.Error():
			code = http.StatusBadRequest

		// User Service
		case User.ErrUserNotFound.Error():
			code = http.StatusNotFound
		case User.ErrInvalidProgramID.Error():
			code = http.StatusBadRequest
		case User.ErrUserExists.Error():
			code = http.StatusBadRequest

		// Unknown error
		default:
			if _, ok := err.(*echo.HTTPError); !ok {
				message = "Internal Error"
			}
		}

		// Send error in a specific format
		controllers.SendErrorResponse(c, code, message)

		// Log errors
		c.Logger().Error(err)
	}
}
