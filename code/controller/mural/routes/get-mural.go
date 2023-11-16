package routes

import (
	"mural/app"
	"mural/config"
	"mural/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetMuaralPage(c echo.Context) error {
	user_key := middleware.GetUserKeyFromContext(c)
	mural_service := c.Get(app.ServiceContextKey).(app.MuralService)
	mural_ses, err := mural_service.DAL.GetMuralForUser(
		user_key,
		mural_service.Config,
	)
	config.Must(err)

	return c.Render(http.StatusOK, "mural.html", mural_ses)
}
