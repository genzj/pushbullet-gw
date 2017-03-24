package server

import "github.com/labstack/echo"

func init() {
	registerEndpoint(func(e *echo.Echo) {
		e.Static("/static", "public/assets")
		e.File("/favicon.ico", "public/assets/images/favicon.ico")
	})
}
