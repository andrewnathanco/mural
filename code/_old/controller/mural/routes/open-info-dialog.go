package routes

import (
	"mural/app"
	"mural/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func OpenInfoDialog(c echo.Context) error {
	user_key := middleware.GetUserKeyFromContext(c)
	mural_service := c.Get(app.ServiceContextKey).(app.MuralService)
	mural_ses, err := mural_service.DAL.GetMuralForUser(
		user_key,
		mural_service.Config,
		false,
	)
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
	}

	return c.Render(http.StatusOK, "info-dialog.html", mural_ses)
}
