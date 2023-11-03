package routes

import (
	"mural/controller/mural/service"
	"mural/db"
	"mural/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func SelectTile(c echo.Context) error {
	i := c.QueryParam("i")
	i_int, err := strconv.ParseInt(i, 10, 64)
    if err != nil {
		return c.String(http.StatusBadRequest, "need to define in the i direction")
    }

	j := c.QueryParam("j")
	j_int, err := strconv.ParseInt(j, 10, 64)
    if err != nil {
		return c.String(http.StatusBadRequest, "need to define in the j direction")
    }

	game_key, err := middleware.GetGameKeyFromContext(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not get game key")
	}

	current_game, err := db.DAL.GetCurrentGame(game_key)
    if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
    }

	new_tiles := service.ResetSelected(current_game.Board.Tiles)
	new_tiles[i_int][j_int].Selected = true

	current_game.Board.Tiles = new_tiles
	current_game.SelectedTile = &new_tiles[i_int][j_int]

	db.DAL.SetCurrentGame(*current_game)

	return c.Render(http.StatusOK, "game-board.html", current_game)
}
