package routes

import (
	"log/slog"
	"mural/controller/mural/service"
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

	user_key := middleware.GetUserKeyFromContext(c)
	curr_mural, err := service.GetCurrentMural(user_key)
	if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "mural-error.html", nil)
	}

	current_tile := curr_mural.Session.Board.Tiles[i_int][j_int]
	if current_tile.Flipped {
		return c.Render(http.StatusOK, "game-board.html", curr_mural.Session)
	}

	curr_mural.Session.Board.Tiles[i_int][j_int].Flipped = true
	curr_mural.Session.Board.Tiles[i_int][j_int].Selected = false
	curr_mural.Session.SelectedTile = nil
	// set game state status
	curr_mural.Session.SessionStatus = model.SESSION_STARTED

	curr_mural.Session.CurrentScore = curr_mural.Session.CurrentScore - current_tile.Penalty
	// db.DAL.SetGameSessionForUser(curr_mural.Session)

	return c.Render(http.StatusOK, "game-board.html", curr_mural)
}
