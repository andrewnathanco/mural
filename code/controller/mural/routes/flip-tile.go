package routes

import (
	"mural/db"
	"mural/middleware"
	"mural/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


func FlipTile(c echo.Context) error {
	i := c.QueryParam("i")
	i_int, err := strconv.ParseInt(i, 10, 64)
    if err != nil {
		c.String(http.StatusBadRequest, "need to define in the i direction")
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

	current_tile := current_game.Board.Tiles[i_int][j_int]
	if current_tile.Flipped {
		return c.Render(http.StatusOK, "game-board.html", current_game)
	}

	current_game.Board.Tiles[i_int][j_int].Flipped = true
	current_game.Board.Tiles[i_int][j_int].Selected = false
	current_game.SelectedTile = nil
	// set game state status
	current_game.GameState = model.GAME_STARTED

	current_game.CurrentScore = current_game.CurrentScore - current_tile.Penalty
	db.DAL.SetCurrentGame(*current_game)

	return c.Render(http.StatusOK, "game-board.html", current_game)
}
