package server

import (
	"errors"
	"html/template"
	"io"
	"os"
	"path/filepath"

	rice "github.com/GeertJohan/go.rice"
	"github.com/labstack/echo"
)

var (
	errNoFileMatchPattern error = errors.New("no matched file found")
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (t *Template) ParseInRice(box *rice.Box, filename string) (*template.Template, error) {
	name := filepath.Base(filename)

	if t.templates == nil {
		t.templates = template.New(name)
	}
	return t.templates.Parse(box.MustString(filename))
}

func (t *Template) ParsePatternInRice(box *rice.Box, dir, pattern string) (*template.Template, error) {
	var err error = nil
	var parsed bool = false

	var walkFn = func(path string, info os.FileInfo, err error) error {
		//fmt.Println("file met:", path, "is embedded?", box.IsEmbedded(), "is appended?", box.IsAppended())
		if info.IsDir() {
			return nil
		}
		if match, err := filepath.Match(pattern, info.Name()); match {
			_, err = t.ParseInRice(box, path)
			if err == nil {
				parsed = true
			}
		}
		return nil
	}
	if _, err := filepath.Match("", pattern); err != nil {
		return nil, err
	}
	walkerr := box.Walk(dir, walkFn)
	if walkerr != nil {
		return nil, walkerr
	} else if err != nil {
		return nil, err
	} else if !parsed {
		return t.templates, errNoFileMatchPattern
	}
	return t.templates, nil
}

var t *Template

func init() {
	registerEndpoint(func(e *echo.Echo) {
		box := rice.MustFindBox("public/views")

		t = &Template{}
		template.Must(t.ParsePatternInRice(box, "/", "*.html"))
		e.Renderer = t
	})
}
