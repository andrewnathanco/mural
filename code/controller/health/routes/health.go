package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthObject struct {

}
func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, HealthObject{})
}
