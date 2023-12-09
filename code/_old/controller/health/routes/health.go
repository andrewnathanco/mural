package routes

import (
	"mural/app"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthStatus string

const (
	DATABSE_FAILED = "FAILED"
	SUCCESS        = "SUCCESS"
)

type HealthObject struct {
	Status HealthStatus
}

func Health(c echo.Context) error {
	mural_service := c.Get(app.ServiceContextKey).(app.MuralService)
	var health HealthObject
	err := mural_service.DAL.PingDatabase()
	if err != nil {
		health.Status = DATABSE_FAILED
	}

	health.Status = SUCCESS
	return c.JSON(http.StatusOK, health)
}
