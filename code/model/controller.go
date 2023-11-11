package model

import (
	"github.com/labstack/echo/v4"
	"github.com/ryanbradynd05/go-tmdb"
)

type IController interface {
	GetTemplates() []TemplateController
	GetRoutes() map[string]func(c echo.Context)error
}

type IRouter interface {
	ConfigureRouter(IController, *echo.Echo)
}

type IMovieController interface {
	GetAnswers(int) ([]tmdb.MovieShort, error)
}