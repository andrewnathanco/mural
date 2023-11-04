package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CloseDialog(c echo.Context) error {
	return c.HTML(http.StatusOK, "")
}
