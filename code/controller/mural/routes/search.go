package routes

import (
	"database/sql"
	"mural/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Search(c echo.Context) error {
    query := c.QueryParam("query")
	if query == "" {
		return c.String(http.StatusOK, "") 
	}

	answers, err := db.DAL.GetAnswersFromQuery(query)
	if err != sql.ErrNoRows {
		if err != nil {
			return c.String(http.StatusInternalServerError, "could not get answers")
		}
	}

	if len(answers) == 0 {
		return c.String(http.StatusOK, "") 
	}

	return c.Render(http.StatusOK, "answer-options.html", answers) 
}
