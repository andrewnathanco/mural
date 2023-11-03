package routes

import (
	"mural/db"
	"mural/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ShareAnswer(c echo.Context) error {
	game_key, err := middleware.GetGameKeyFromContext(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not get game key")
	}

	_, err = db.DAL.GetCurrentGame(game_key)

    if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
    }
    if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
    }

	return c.String(http.StatusOK, "shared")
}
