package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func init() {
	registerEndpoint(func(e *echo.Echo) {
		e.GET("/welcome/:secret", welcome)
	})
}

func welcome(c echo.Context) error {
	cc := c.(*pushbulletClientContext)
	secret := c.Param("secret")
	user, err := cc.backend.GetBySecret(secret, false)
	if err != nil {
		return c.Render(http.StatusForbidden, "error", err.Error())
	}
	cc.client.LoadToken(user.Tokens[0].AccessToken)
	if me, err := cc.client.GetUser(); err != nil {
		return c.Render(http.StatusOK, "error", err.Error())
	} else {
		return c.Render(http.StatusOK, "hello", me.Name)
	}
}
