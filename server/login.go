package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func init() {
	registerEndpoint(func(e *echo.Echo) {
		e.GET("/login", login)
		e.GET("/", login)
	})
}

func login(c echo.Context) error {
	cc := c.(*pushbulletClientContext)
	return c.Render(http.StatusOK, "login", cc.client.AuthURL())
}
