package router

import (
	"net/http"

	"github.com/epointpayment/key_management_system/app/controllers"

	"github.com/labstack/echo"
)

func (r *Router) appendErrorHandler() {
	r.e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		message := err.Error()
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
			message = he.Message.(string)
		}

		switch message {
		}

		controllers.SendErrorResponse(c, code, message)
		c.Logger().Error(err)
	}
}
