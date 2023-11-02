package mural

import (
	"html/template"
	"mural/db"
	"mural/model"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	board_size = 6
)

type MuralController struct {

}

func mod(a, b int) int {
    return a % b
}


func NewMuralController() (*MuralController) {
	board := newGameBoard(board_size)


	current_game := model.Game{
		CurrentScore: 56,
		Board: *board,
		Poster: model.Poster{
			Name: "Talk To Me",
			URL: "https://image.tmdb.org/t/p/w1280//kdPMUMJzyYAc4roD52qavX0nLIC.jpg",
		},
	}

	db.DAL.SetCurrentGame(current_game)
	
	return &MuralController{}
}

// GetTemplates() []model.TemplateController
// GetRoutes() map[string]func(c echo.Context)error

func (mc MuralController) GetTemplates() []model.TemplateController {
	mod_map := template.FuncMap{
        "mod": mod,
    }

	mural_template := template.Must(template.New("mural_template").Funcs(mod_map).ParseFiles("view/mural/mural.html", "view/mural/game/game-board.html"))
	mural_template_controller := model.TemplateController{
		Template: mural_template,
		Name: "mural.html",
	}

	error_template := template.Must(template.ParseFiles("view/mural/mural-error.html"))
	error_template_controller := model.TemplateController{
		Template: error_template,
		Name: "mural-error.html",
	}
	templates := []model.TemplateController{
		mural_template_controller, error_template_controller,
	}
	return templates
}

func (mc MuralController) GetRoutes() map[string]func(c echo.Context) error {
	router := map[string]func(c echo.Context) error{}
	router["flip-tile"] = flipTile
	router["mural"] = getMuaralPage
	return router
}

func flipTile(c echo.Context) error {
	i := c.QueryParam("i")
	i_int, err := strconv.ParseInt(i, 10, 64)
    if err != nil {
		c.String(http.StatusBadRequest, "need to define in the i direction")
    }

	j := c.QueryParam("j")
	j_int, err := strconv.ParseInt(j, 10, 64)
    if err != nil {
		c.String(http.StatusBadRequest, "need to define in the j direction")
    }

	current_game, err := db.DAL.GetCurrentGame()
    if err != nil {
		c.String(http.StatusInternalServerError, "could not get current game")
    }

	current_game.Board.Tiles[i_int][j_int].Flipped = true
	db.DAL.SetCurrentGame(*current_game)

	dat, _ := os.ReadFile("/view/mural/game/flipped_tile.html")
	return c.String(http.StatusOK, string(dat))
}

func getMuaralPage(c echo.Context) error {
	current_game, err := db.DAL.GetCurrentGame()
	if err != nil {
		return c.Render(http.StatusInternalServerError, "mural-error.html", nil)
	}
	return c.Render(http.StatusOK, "mural.html", current_game)
}
