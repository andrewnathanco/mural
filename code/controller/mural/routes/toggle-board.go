package routes

import (
	"mural/app"
	"mural/config"
	"mural/db"
	"mural/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ToggleBoard(c echo.Context) error {
	board_state := c.QueryParam("state")
	user_key := middleware.GetUserKeyFromContext(c)
	mural_service := c.Get(app.ServiceContextKey).(app.MuralService)
	mural_ses, err := mural_service.DAL.GetMuralForUser(
		user_key,
		mural_service.Config,
	)

	if board_state == db.BOARD_AS_GAME {
		mural_ses.BoardState = db.BOARD_AS_GAME
	} else {
		mural_ses.BoardState = db.BOARD_NORMAL
	}

	config.Must(err)
	return c.Render(http.StatusOK, "game-board.html", mural_ses)
}
