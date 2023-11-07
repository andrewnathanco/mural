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

func SelectAnswer(c echo.Context) error {
	id := c.QueryParam("answer")
	if id == "" {
		return c.String(http.StatusBadRequest, "need to define an id")
	}

	id_int, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
	}

	user_key := middleware.GetUserKeyFromContext(c)
	curr_mural, err := service.GetCurrentMural(user_key)
    if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
    }


	var selected_answer model.Answer
	for _, answer := range curr_mural.Game.Answers {
		if answer.ID == int(id_int) {
			selected_answer = answer
		}
	}

	curr_mural.Session.SelectedAnswer = &selected_answer
	db.DAL.SetGameSessionForUser(curr_mural.Session)
	return c.Render(http.StatusOK, "answers.html", curr_mural)
}
