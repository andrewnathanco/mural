package routes

import (
	"mural/controller/mural/service"
	"mural/db"
	"mural/middleware"
	"mural/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)


func SetHardMode(c echo.Context) error {
	user_key := middleware.GetUserKeyFromContext(c)
	curr_mural, err := service.GetCurrentMural(user_key)
    if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
    }

	if curr_mural.Session.SessionStatus == model.SESSION_STARTED || curr_mural.Session.SessionStatus == model.SESSION_OVER {
		return c.Render(http.StatusOK, "game-board.html", curr_mural)
	}

	enabled := c.QueryParam("enabled")
	hard_mode_enabled, _ := strconv.ParseBool(enabled)
	curr_mural.UserData.HardModeEnabled = hard_mode_enabled

	db.DAL.SetUserData(user_key, curr_mural.UserData)
	return c.Render(http.StatusOK, "game-board.html", curr_mural)
}
