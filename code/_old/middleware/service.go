package middleware

import (
	"mural/app"

	"github.com/labstack/echo/v4"
)

// Process is the middleware function.
func PassServiceData(service app.MuralService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(app.ServiceContextKey, service)
			return next(c)
		}
	}
}
