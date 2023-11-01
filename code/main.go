package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"mural/controller/mural"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	ErrCouldNotParseTempaltes = fmt.Errorf("could not parse templates")
)

type TemplateRenderer struct {
	templates map[string]*template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]

	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}

	return tmpl.ExecuteTemplate(w, name, data)
}

func main() {
	echo.NotFoundHandler = func(c echo.Context) error {
		return c.Render(http.StatusNotFound, "404.html", nil)
	}

	e := echo.New()

	// define templates
	templates := map[string]*template.Template{}


	// setup controllers
	mural_template, handle_mural_route, err := mural.NewMuralController() 
	if err != nil {
		e.Logger.Fatal(err)
	}

	// add templates
	templates[mural_template.Name] = mural_template.Template
	e.Renderer = &TemplateRenderer{
		templates: templates,
	}

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "./static")

	e.GET("/", handle_mural_route)

	e.Logger.Fatal(e.Start(":1323"))
}
