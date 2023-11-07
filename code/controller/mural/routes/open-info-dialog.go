package routes

import (
	"mural/controller/mural/service"
	"mural/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func OpenInfoDialog(c echo.Context) error {
	user_key := middleware.GetUserKeyFromContext(c)
	curr_mural, err := service.GetCurrentMural(user_key)
    if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
    }

	return c.Render(http.StatusOK, "info-dialog.html", curr_mural)
}
