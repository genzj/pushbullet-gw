package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func init() {
	registerEndpoint(func(e *echo.Echo) {
		e.GET("/welcome", welcome)
	})
}

func welcome(c echo.Context) error {
	cc := c.(*pushbulletClientContext)
	if me, err := cc.client.GetUser(); err != nil {
		return c.Render(http.StatusOK, "error", err.Error())
	} else {
		return c.Render(http.StatusOK, "hello", me.Name)
	}
}
