package routes

import (
	"mural/controller/mural/service"
	"mural/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func OpenStatsDialog(c echo.Context) error {
	user_key := middleware.GetUserKeyFromContext(c)
	curr_mural, err := service.GetCurrentMural(user_key)
    if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
    }

	return c.Render(http.StatusOK, "stats-dialog.html", curr_mural)
}
