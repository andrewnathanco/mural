package routes

import (
	"fmt"
	"log/slog"
	"mural/app"
	"mural/db"
	"mural/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ShareLink struct {
	Mural db.Mural
	Link  string
}

func CreateShareLink(c echo.Context) error {
	user_key := middleware.GetUserKeyFromContext(c)
	name := c.FormValue("name")
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

	mural_ses.User.DisplayName = name
	err = mural_service.DAL.UpsertUser(mural_ses.User)
	if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "mural-error.html", nil)
	}

	link := fmt.Sprintf("%s/share?user_key=%s", mural_service.Config.AppURL, user_key)
	return c.Render(http.StatusOK, "share-link.html", ShareLink{Link: link, Mural: mural_ses})
}
