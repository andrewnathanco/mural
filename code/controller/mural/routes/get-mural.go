package routes

import (
	"log/slog"
	"mural/controller/mural/service"
	"mural/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetMuaralPage(c echo.Context) error {
	user_key := middleware.GetUserKeyFromContext(c)
	curr_mural, err := service.GetCurrentMural(user_key)
	if err != nil {
		slog.Error(err.Error())
		return c.Render(http.StatusInternalServerError, "mural-error.html", nil)
	}

	return c.Render(http.StatusOK, "mural.html", curr_mural)
}