package routes

import (
	"fmt"
	"log/slog"
	"mural/app"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetShare(c echo.Context) error {
	mural_service := c.Get(app.ServiceContextKey).(app.MuralService)
	user_key := c.QueryParam("user-key")

	user_exists, err := mural_service.DAL.CheckForUser(user_key)
	if err != nil || !user_exists {
		slog.Error(fmt.Errorf("no user provided").Error())
		return c.Render(http.StatusInternalServerError, "share-error.html", nil)
	}

	mural_ses, err := mural_service.DAL.GetMuralForUser(
		user_key,
		mural_service.Config,
		true,
	)

	if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "share-error.html", nil)
	}

	return c.Render(http.StatusOK, "share.html", mural_ses)
}
