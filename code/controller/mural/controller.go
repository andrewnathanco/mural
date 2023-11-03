package mural

import (
	"html/template"
	"mural/controller/shared"
	"mural/db"
	"mural/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


type MuralController struct { }

// functions
func mod(a, b int) int {
    return a % b
}

type FlipButton struct {
	Button shared.Button
	Tile *model.Tile
}

func newFlipButton(
	text string,
	disabled bool,
	tile *model.Tile,
) FlipButton {
	return FlipButton{
		Button: shared.Button{
			Text: text,
			Disabled: disabled,
		},
		Tile: tile,
	}
}

func NewMuralController() (*MuralController) {
	board := newGameBoard(board_size)
	correct_movie, answers := newAnswers()

	current_game := model.Game{
		CurrentScore: 56,
		Board: *board,
		Current: correct_movie,
		Answers: answers,
	}

	db.DAL.SetCurrentGame(current_game)
	
	return &MuralController{}
}

// GetTemplates() []model.TemplateController
// GetRoutes() map[string]func(c echo.Context)error

func (mc MuralController) GetTemplates() []model.TemplateController {
	func_map := template.FuncMap{
        "mod": mod,
		"newButton": shared.NewButton,
		"newFlipButton": newFlipButton,
    }

	mural_template := template.Must(
		template.New("mural_template").Funcs(func_map).ParseFiles(
			"view/mural/mural.html", 
			"view/mural/mural.tmpl", 
			"view/mural/game/board/game-board.html",
			"view/mural/game/board/game-board.tmpl",
			"view/mural/game/answers.html",
			"view/mural/game/answer/default-answer.html",
			"view/mural/game/tile/default-tile.html",
			"view/mural/game/tile/selected/selected-tile.tmpl",
			"view/mural/game/tile/flipped-tile.html",
			"view/mural/buttons/flip-button.tmpl",
			"view/mural/buttons/submit-button.tmpl",
		))



	mural_template_controller := model.TemplateController{
		Template: mural_template,
		Name: "mural.html",
	}

	error_template := template.Must(
		template.New("mural-error").ParseFiles("view/mural/mural-error.html"),
	)
	error_template_controller := model.TemplateController{
		Template: error_template,
		Name: "mural-error.html",
	}

	selected_tile := template.Must(
		template.New("selected").Funcs(func_map).
		ParseFiles(
			"view/mural/game/board/game-board.html",
			"view/mural/game/board/game-board.tmpl",
			"view/mural/game/tile/default-tile.html",
			"view/mural/game/tile/selected/selected-tile.tmpl",
			"view/mural/game/tile/flipped-tile.html",
			"view/mural/buttons/flip-button.tmpl",
			"view/mural/buttons/submit-button.tmpl",
		),
	)

	selected_tile_controller := model.TemplateController{
		Template: selected_tile,
		Name: "game-board.html",
	}

	flipped_tile := template.Must(
		template.New("flipped").Funcs(func_map).
		ParseFiles(
			"view/mural/game/board/game-board.html",
			"view/mural/game/board/game-board.tmpl",
			"view/mural/game/tile/default-tile.html",
			"view/mural/game/tile/selected/selected-tile.tmpl",
			"view/mural/game/tile/flipped-tile.html",
			"view/mural/buttons/flip-button.tmpl",
		),
	)

	flipped_tile_controller := model.TemplateController{
		Template: flipped_tile,
		Name: "game-board.html",
	}

	templates := []model.TemplateController{
		mural_template_controller, error_template_controller,
		selected_tile_controller, flipped_tile_controller,
	}
	return templates
}

func (mc MuralController) GetRoutes() map[string]func(c echo.Context) error {
	router := map[string]func(c echo.Context) error{}
	router["flip-tile"] = flipTile
	router["mural"] = getMuaralPage
	router["select-tile"] = selectTile
	return router
}

func selectTile(c echo.Context) error {
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

	new_tiles := resetSelected(current_game.Board.Tiles)
	new_tiles[i_int][j_int].Selected = true

	current_game.Board.Tiles = new_tiles
	current_game.Selected = &new_tiles[i_int][j_int]

	db.DAL.SetCurrentGame(*current_game)

	return c.Render(http.StatusOK, "game-board.html", current_game)
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

	current_tile := current_game.Board.Tiles[i_int][j_int]
	if current_tile.Flipped {
		return c.Render(http.StatusOK, "game-board.html", current_game)
	}

	current_game.Board.Tiles[i_int][j_int].Flipped = true
	current_game.Board.Tiles[i_int][j_int].Selected = false
	current_game.Selected = nil

	current_game.CurrentScore = current_game.CurrentScore - current_tile.Penalty
	db.DAL.SetCurrentGame(*current_game)

	return c.Render(http.StatusOK, "game-board.html", current_game)
}



func getMuaralPage(c echo.Context) error {
	current_game, err := db.DAL.GetCurrentGame()
	if err != nil {
		return c.Render(http.StatusInternalServerError, "mural-error.html", nil)
	}
	return c.Render(http.StatusOK, "mural.html", current_game)
}
