package routes

import (
	"log/slog"
	"mural/controller/mural/service"
	"mural/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func OpenCopiedAlert(c echo.Context) error {
	alert := c.QueryParam("alert")
	if alert == "" {
		return c.String(http.StatusBadRequest, "need to declare an alert")
	}

	user_key, err := middleware.GetUserKeyFromContext(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not get game key")
	}

	curr_mural, err := service.GetCurrentMural(user_key)
	if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "mural-error.html", nil)
	}

	if alert == "success" {
		return c.Render(http.StatusOK, "copied-success.html", curr_mural)
	} else {
		return c.Render(http.StatusOK, "copied-error.html", curr_mural)
	}
}
