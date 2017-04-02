package server

import (
	"net/http"

	"github.com/genzj/pushbullet-gw/pushbullet"
	"github.com/labstack/echo"
)

func init() {
	registerEndpoint(func(e *echo.Echo) {
		e.GET("/welcome/:secret", welcome)
	})
}

func welcome(c echo.Context) error {
	type datum struct {
		Name    string
		Devices []pushbullet.Device
	}
	cc := c.(*pushbulletClientContext)
	secret := c.Param("secret")
	user, err := cc.backend.GetBySecret(c, secret, false)
	if err != nil {
		return c.Render(http.StatusForbidden, "error", err.Error())
	}
	cc.client.LoadToken(user.Tokens[0].AccessToken)
	if me, err := cc.client.GetUser(); err != nil {
		return c.Render(http.StatusOK, "error", err.Error())
	} else if devices, err := cc.client.ListDevices(); err != nil {
		return c.Render(http.StatusOK, "error", err.Error())
	} else {
		err := c.Render(http.StatusOK, "hello", &datum{
			Name:    me.Name,
			Devices: devices.Devices,
		})
		if err != nil {
			return c.Render(http.StatusOK, "error", err.Error())
		} else {
			return nil
		}
	}
}
