package mural

import (
	"mural/model"

	"github.com/labstack/echo/v4"
)

type MuralRouter struct {

}

func NewMuralRouter() MuralRouter{
	return MuralRouter{}
}

func (r MuralRouter) ConfigureRouter(c model.IController, e *echo.Echo) {
	e.GET("/", c.GetRoutes()["mural"])
	e.PUT("/mural/flip-tile", c.GetRoutes()["flip-tile"])
	e.PUT("/mural/select-tile", c.GetRoutes()["select-tile"])
}