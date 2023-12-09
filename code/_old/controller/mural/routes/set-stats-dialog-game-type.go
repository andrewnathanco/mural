package routes

import (
	"mural/app"
	"mural/db"
	"mural/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetStatsDialogGameType(c echo.Context) error {
	game_type := c.QueryParam("type")
	if game_type != db.REGULAR_MODE && game_type != db.EASY_MODE {
		game_type = db.REGULAR_MODE
	}
	user_key := middleware.GetUserKeyFromContext(c)
	mural_service := c.Get(app.ServiceContextKey).(app.MuralService)
	mural_ses, err := mural_service.DAL.GetMuralForUser(
		user_key,
		mural_service.Config,
		false,
	)

	if err != nil {
		return c.String(http.StatusInternalServerError, "couldn't get mural")
	}

	return c.Render(http.StatusOK, "stats-dialog.html", StatsState{
		Mural:    mural_ses,
		GameType: game_type,
	})
}
