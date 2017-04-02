package server

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var t *Template

func init() {
	registerEndpoint(func(e *echo.Echo) {
		t = &Template{
			// TODO embed template files with go.rice
			templates: template.Must(template.ParseGlob("server/public/views/*.html")),
		}
		e.Renderer = t
	})
}

//func init() {
//t = &Template{
//templates: template.Must(template.ParseGlob("public/views/*.html")),
//}
//e.Renderer = t
//}
