package controller

import (
	"errors"
	"html/template"
	"io"
	"mural/controller/health"
	"mural/controller/mural"
	"mural/model"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	Templates map[string]*template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.Templates[name]

	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}

	return tmpl.ExecuteTemplate(w, name, data)
}



type RouteController struct {
	Router model.IRouter
	Controller model.IController
}

func GetRouteControllers() ([]RouteController) {
	route_controllers := []RouteController{
		{
			Router: mural.NewMuralRouter(),
			Controller: mural.NewMuralController(),
		},
		{
			Router: health.NewHealthRouter(),
			Controller: health.NewHealthController(),
		},
	}

	return route_controllers
}