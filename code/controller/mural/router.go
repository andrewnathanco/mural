package mural

import (
	"mural/controller/mural/routes"
	"mural/model"

	"github.com/labstack/echo/v4"
)

type MuralRouter struct {
}

func NewMuralRouter() MuralRouter {
	return MuralRouter{}
}

func (mc MuralController) GetRoutes() map[string]func(c echo.Context) error {
	router := map[string]func(c echo.Context) error{}
	router["flip-tile"] = routes.FlipTile
	router["mural"] = routes.GetMuaralPage
	router["select-tile"] = routes.SelectTile
	router["select-option"] = routes.SelectOption
	router["submit"] = routes.Submit
	router["set-hard-mode"] = routes.SetHardMode
	router["open-stats-dialog"] = routes.OpenStatsDialog
	router["set-stats-dialog-game-type"] = routes.SetStatsDialogGameType
	router["open-info-dialog"] = routes.OpenInfoDialog
	router["open-share-dialog"] = routes.OpenShareDialog
	router["toggle-board-state"] = routes.ToggleBoard
	router["search"] = routes.Search
	router["share"] = routes.GetShare
	router["create-share-link"] = routes.CreateShareLink
	return router
}

func (r MuralRouter) ConfigureRouter(c model.IController, e *echo.Echo) {
	e.GET("/", c.GetRoutes()["mural"])
	e.GET("/share", c.GetRoutes()["share"])
	e.PUT("/mural/flip-tile", c.GetRoutes()["flip-tile"])
	e.PUT("/mural/select-tile", c.GetRoutes()["select-tile"])
	e.PUT("/mural/select-option", c.GetRoutes()["select-option"])
	e.PUT("/mural/submit", c.GetRoutes()["submit"])
	e.PUT("/mural/set-hard-mode", c.GetRoutes()["set-hard-mode"])
	e.PUT("/mural/toggle-board-state", c.GetRoutes()["toggle-board-state"])
	e.POST("/mural/search", c.GetRoutes()["search"])
	e.GET("/mural/open-stats-dialog", c.GetRoutes()["open-stats-dialog"])
	e.GET("/mural/open-share-dialog", c.GetRoutes()["open-share-dialog"])
	e.PUT("/mural/set-stats-dialog-game-type", c.GetRoutes()["set-stats-dialog-game-type"])
	e.GET("/mural/open-info-dialog", c.GetRoutes()["open-info-dialog"])
	e.PUT("/mural/create-share-link", c.GetRoutes()["create-share-link"])
}
