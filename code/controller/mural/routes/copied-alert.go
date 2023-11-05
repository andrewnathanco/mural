package routes

import (
	"mural/db"
	"mural/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func OpenCopiedAlert(c echo.Context) error {
	alert := c.QueryParam("alert")
	if alert == "" {
		return c.String(http.StatusBadRequest, "need to declare an alert")
	}

	game_key, err := middleware.GetGameKeyFromContext(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not get game key")
	}

	game, err := db.DAL.GetCurrentGame(game_key)

    if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
    }
    if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
    }

	if alert == "success" {
		return c.Render(http.StatusOK, "copied-success.html", game)
	} else {
		return c.Render(http.StatusOK, "copied-error.html", game)
	}
}
