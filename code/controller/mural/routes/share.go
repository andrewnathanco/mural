package routes

import (
	"mural/api"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Share(c echo.Context) error {
	// do analytics stuff
	api.AnalyticsController.RegisterEvent(api.EVENT_SHARE, c.Request())
	return c.String(http.StatusOK, "")
}
