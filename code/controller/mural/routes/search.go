package routes

import (
	"database/sql"
	"mural/app"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Search(c echo.Context) error {
	query := c.FormValue("search-query")
	mural_service := c.Get(app.ServiceContextKey).(app.MuralService)

	options, err := mural_service.DAL.GetOptionsByQuery(query)
	if err != sql.ErrNoRows {
		if err != nil {
			return c.String(http.StatusInternalServerError, "could not get answers")
		}
	}

	if len(options) == 0 {
		return c.String(http.StatusOK, "")
	}

	return c.Render(http.StatusOK, "answer-options.html", options)
}
