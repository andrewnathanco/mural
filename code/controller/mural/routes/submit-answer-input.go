package routes

import (
	"fmt"
	"mural/controller/mural/service"
	"mural/middleware"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/lithammer/fuzzysearch/fuzzy"
)


func SubmitAnswerInput(c echo.Context) error {
	user_key := middleware.GetUserKeyFromContext(c)
	curr_mural, err := service.GetCurrentMural(user_key)
    if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
    }

	answer := c.QueryParam("answer")
	if answer == "" {
		return c.Render(http.StatusBadRequest, "need-answer.html", nil)
		// return dialog
	}

	// lets fuzzy check
	number := fuzzy.RankMatch(answer, curr_mural.Game.CorrectAnswer.Name)
	fmt.Println(number)

	return c.Render(http.StatusOK, "game-board.html", curr_mural)
}
