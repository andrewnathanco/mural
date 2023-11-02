package controller

import (
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
	}


	return route_controllers
}