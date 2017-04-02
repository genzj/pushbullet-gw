package server

import (
	"net/http"

	"github.com/labstack/echo"
)

func init() {
	registerEndpoint(func(e *echo.Echo) {
		e.GET("/push/:secret/:deviceID", push)
	})
}

func push(c echo.Context) error {
	cc := c.(*pushbulletClientContext)
	secret := c.Param("secret")
	deviceID := c.Param("deviceID")
	user, err := cc.backend.GetBySecret(c, secret, false)
	if err != nil {
		// TODO add a redirection page to guide re-autentication
		return c.Render(http.StatusForbidden, "error", err.Error())
	}
	cc.client.LoadToken(user.Tokens[0].AccessToken)
	if res, err := cc.client.PushLink(deviceID, cc.QueryParam("t"), cc.QueryParam("s"), cc.QueryParam("u")); err != nil {
		return c.Render(http.StatusOK, "error", err.Error())
	} else {
		return c.JSON(http.StatusOK, res)
	}
}
