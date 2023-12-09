package routes

import (
	"mural/app"
	"mural/db"
	"mural/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetHardMode(c echo.Context) error {
	user_key := middleware.GetUserKeyFromContext(c)
	mural_service := c.Get(app.ServiceContextKey).(app.MuralService)

	mode := c.QueryParam("mode")
	user := db.User{
		UserKey: user_key,
	}

	if mode == db.EASY_MODE {
		user.GameType = db.EASY_MODE
	} else {
		user.GameType = db.REGULAR_MODE
	}
	err := mural_service.DAL.UpsertUser(user)
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
	}

	mural_ses, err := mural_service.DAL.GetMuralForUser(
		user_key,
		mural_service.Config,
		false,
	)
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
	}

	return c.Render(http.StatusOK, "game-board.html", mural_ses)
}
