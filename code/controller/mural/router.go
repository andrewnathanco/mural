package mural

import (
	"fmt"
	"mural/model"

	"github.com/labstack/echo/v4"
)

type MuralRouter struct {

}

func NewMuralRouter() MuralRouter{
	return MuralRouter{}
}

func (r MuralRouter) ConfigureRouter(c model.IController, e *echo.Echo) {
	fmt.Println(c.GetRoutes())
	e.GET("/", c.GetRoutes()["mural"])
	e.PUT("/mural/flip", c.GetRoutes()["flip-tile"])
}