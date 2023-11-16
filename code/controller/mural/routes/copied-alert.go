package routes

import (
	"mural/app"
	"mural/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func OpenCopiedAlert(c echo.Context) error {
	alert := c.QueryParam("alert")
	if alert == "" {
		return c.String(http.StatusBadRequest, "need to declare an alert")
	}

	user_key := middleware.GetUserKeyFromContext(c)
	mural_service := c.Get(app.ServiceContextKey).(app.MuralService)
	mural_ses, err := mural_service.DAL.GetMuralForUser(
		user_key,
		mural_service.Config,
	)

	if err != nil {
		return c.String(http.StatusInternalServerError, "couldn't get mural")
	}

	if alert == "success" {
		return c.Render(http.StatusOK, "copied-success.html", mural_ses)
	} else {
		return c.Render(http.StatusOK, "copied-failure.html", mural_ses)
	}
}
