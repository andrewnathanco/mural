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
	answer_param := c.QueryParam("answer")
	if answer_param == "" {
		return c.String(http.StatusBadRequest, "need to define an id")
	}

	user_key := middleware.GetUserKeyFromContext(c)
	curr_mural, err := service.GetCurrentMural(user_key)
    if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
    }

	if curr_mural.UserData.HardModeEnabled {
		answer, err := db.DAL.GetAnswerFromKey(answer_param)
		if err != nil {
			return c.String(http.StatusInternalServerError, "could not get current game")
		}

		curr_mural.Session.SelectedAnswer = answer
		db.DAL.SetGameSessionForUser(curr_mural.Session)
		return c.Render(http.StatusOK, "answer-input.html", curr_mural)
	}

	id_int, err := strconv.ParseInt(answer_param, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "invalid id")
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
