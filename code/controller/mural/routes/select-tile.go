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

	user_key, err := middleware.GetUserKeyFromContext(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not get game key")
	}

	curr_mural, err := service.GetCurrentMural(user_key)
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not get game key")
	}

	new_tiles := service.ResetSelected(curr_mural.Session.Board.Tiles)
	new_tiles[i_int][j_int].Selected = true

	curr_mural.Session.Board.Tiles = new_tiles
	curr_mural.Session.SelectedTile = &new_tiles[i_int][j_int]

	db.DAL.SetGameSessionForUser(curr_mural.Session)

	return c.Render(http.StatusOK, "game-board.html", curr_mural)
}
