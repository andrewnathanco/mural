package mural

import (
	"html/template"
	"mural/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Poster struct {
	Name string
	URL string
}

type Game struct {
	CurrentScore int
	GameBoard [][]bool
	Poster Poster
}



func NewMuralController() (*controller.TemplateController, func(c echo.Context) error, error) {
	mural_template := template.Must(template.ParseFiles("view/mural/mural.html", "view/mural/game/game-board.html"))
	new_template_controller := controller.TemplateController{
		Template: mural_template,
		Name: "mural.html",
	}
	
	return &new_template_controller, GetMural, nil
}

func GetMural(c echo.Context) error {
	// TODO: replace with database
	game := Game{
		CurrentScore: 56,
		GameBoard: [][]bool{
			{
				true, true, true, true, true, false,
			},
			{
				true, true, false, true, true, true,
			},
			{
				true, false, true, true, true, true,
			},
			{
				true, true, true, false, true, false,
			},
			{
				true, true, true, true, true, true,
			},
			{
				true, false, true, true, true, false,
			},
		},
		Poster: Poster{
			Name: "Talk To Me",
			URL: "https://image.tmdb.org/t/p/w1280//kdPMUMJzyYAc4roD52qavX0nLIC.jpg",
		},
	}

	return c.Render(http.StatusOK, "mural.html", game)
}
