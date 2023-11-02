package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"mural/controller"
	"mural/db"
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
	// setup database
	db.DAL = db.NewMemoryDAL()


	echo.NotFoundHandler = func(c echo.Context) error {
		return c.Render(http.StatusNotFound, "404.html", nil)
	}

	e := echo.New()

	// define templates
	templates := map[string]*template.Template{}


	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// setup routes and controllers
	route_conrollers := controller.GetRouteControllers()


	for _, route_controller := range route_conrollers {
		// add templates
		for _, template := range route_controller.Controller.GetTemplates() {
			templates[template.Name] = template.Template
		}

		// add routes
		route_controller.Router.ConfigureRouter(route_controller.Controller, e)
	}


	e.Renderer = &TemplateRenderer{
		templates: templates,
	}

	// setup routes
	e.Static("/static", "./static")
	e.Logger.Fatal(e.Start(":1323"))
}
