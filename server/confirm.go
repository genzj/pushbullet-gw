package server

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/genzj/pushbullet-gw/storage"
	"github.com/labstack/echo"
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
	}

	if me, err := cc.client.GetUser(); err != nil {
		panic(err.Error())
	} else {
		user, err := cc.backend.GetByPushbulletID(me.ID)
		if err != nil {
			if user, err = cc.backend.NewUser(&storage.User{
				PushbulletID: me.ID,
				Name:         me.Name,
				Email:        me.Email,
			}); err != nil {
				panic(err)
			}
		}

		if _, err := cc.backend.IssueToken(user, cc.client.TokenToSave()); err != nil {
			panic(err)
		}
		return c.Redirect(http.StatusTemporaryRedirect, "/welcome/"+user.SimplePushSecret)
	}
}
