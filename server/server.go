package server

import (
	log "github.com/Sirupsen/logrus"
	"github.com/genzj/pushbullet-gw/pushbullet"
	"github.com/genzj/pushbullet-gw/storage"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Server setup a gateway for simpler access to Pushbullet APIs
// for services only support HTTP access w/o OAuth2 such as innoreader

type register func(*echo.Echo)

type pushbulletClientContext struct {
	echo.Context
	client  *pushbullet.Client
	backend storage.Backend
}

var ends []register

func registerEndpoint(end register) {
	ends = append(ends, end)
}

func newEcho(client *pushbullet.Client, backend storage.Backend) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// inject client instance by copy the passed in instance
			cc := &pushbulletClientContext{
				Context: c,
				client:  client.SafeClone(),
				backend: backend,
			}
			return h(cc)
		}
	})

	for _, reg := range ends {
		reg(e)
	}
	return e
}

func Start(client *pushbullet.Client, backend storage.Backend, listen string) {
	e := newEcho(client, backend)
	log.Info("listening on ", listen)
	e.Logger.Fatal(e.Start(listen))
}
