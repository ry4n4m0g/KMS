package router

import (
	"github.com/epointpayment/key_management_system/app/router/middleware/auth"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func (r *Router) appendMiddleware() {
	r.e.Use(middleware.Gzip())
	r.e.Use(middleware.Logger())
	r.e.Use(middleware.Recover())
}

func (r *Router) mwBasicAuth() echo.MiddlewareFunc {
	return middleware.BasicAuth(auth.BasicValidator)
}
