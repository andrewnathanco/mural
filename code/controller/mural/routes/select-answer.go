package routes

import (
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

	game_key, err := middleware.GetGameKeyFromContext(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not get game key")
	}

	current_game, err := db.DAL.GetCurrentGame(game_key)
    if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
    }

	answers := []model.Answer{}
	selected_answer := model.Answer{}
	for _, a := range current_game.Answers {
		answer := model.Answer{
			Movie: a.Movie,
			IsCorrect: a.IsCorrect,
			Selected: false,
		}

		if int(id_int) == a.Movie.ID {
			answer.Selected = true
			selected_answer = answer
		}


		answers = append(answers, answer)
	}

	current_game.SelectedAnswer = &selected_answer
	current_game.Answers = answers
	db.DAL.SetCurrentGame(*current_game)
	return c.Render(http.StatusOK, "answers.html", current_game)
}
