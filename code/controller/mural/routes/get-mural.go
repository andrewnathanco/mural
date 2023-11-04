package routes

import (
	"log/slog"
	"mural/db"
	"mural/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetMuaralPage(c echo.Context) error {
	game_key, err := middleware.GetGameKeyFromContext(c)
	if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "mural-error.html", nil)
	}

	current_game, err := db.DAL.GetCurrentGame(game_key)
    if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "mural-error.html", nil)
    }

	return c.Render(http.StatusOK, "mural.html", current_game)
}