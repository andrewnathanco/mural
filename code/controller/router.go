package controller

import (
	"mural/controller/health"
	"mural/controller/mural"
	"mural/model"
)

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