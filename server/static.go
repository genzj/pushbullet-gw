package server

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
)

func init() {
	registerEndpoint(func(e *echo.Echo) {
		// the file server for rice. "public/assets" is the folder where the files come from.
		box := rice.MustFindBox("public/assets")
		assetHandler := http.FileServer(box.HTTPBox())
		e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))
		e.GET("/favicon.ico", func(c echo.Context) error {
			return c.Blob(http.StatusOK, "image/x-icon", box.MustBytes("images/favicon.ico"))
		})
	})
}
