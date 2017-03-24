package server

import (
	log "github.com/Sirupsen/logrus"
	"github.com/labstack/echo"
	"net/http"
)

func init() {
	registerEndpoint(func(e *echo.Echo) {
		e.GET("/auth_confirm", confirm)
	})
}

func confirm(c echo.Context) error {
	cc := c.(*pushbulletClientContext)
	client := cc.client
	client.Credential.Code = c.QueryParam("code")
	log.Debug("Code received: ", c.QueryParam("code"))

	if err := client.RefreshToken(); err != nil {
		return c.Render(http.StatusOK, "error", err.Error())
	} else {
		return c.Redirect(http.StatusTemporaryRedirect, "/welcome")
	}
}
