package routes

import (
	"log/slog"
	"mural/app"
	"mural/db"
	"mural/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ToggleBoard(c echo.Context) error {
	board_state := c.QueryParam("state")
	read_only := c.QueryParam("read_only")
	user_key := middleware.GetUserKeyFromContext(c)
	mural_service := c.Get(app.ServiceContextKey).(app.MuralService)
	mural_ses, err := mural_service.DAL.GetMuralForUser(
		user_key,
		mural_service.Config,
		false,
	)

	if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "mural-error.html", nil)
	}

	if board_state == db.BOARD_AS_GAME {
		mural_ses.BoardState = db.BOARD_AS_GAME
	} else {
		mural_ses.BoardState = db.BOARD_NORMAL
	}

	read_only_bool, _ := strconv.ParseBool(read_only)
	mural_ses.ReadOnly = read_only_bool

	return c.Render(http.StatusOK, "game-board.html", mural_ses)
}
