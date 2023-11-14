package routes

import (
	"mural/controller/mural/service"
	"mural/middleware"
	"mural/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AnswerType string

func SubmitAnswer(c echo.Context) error {
	option := c.QueryParam("type")

	user_key := middleware.GetUserKeyFromContext(c)
	curr_mural, err := service.GetCurrentMural(user_key)
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
	}

	curr_mural.Session.GameWon = true
	if option == "give-up" {
		curr_mural.Session.CurrentScore = 0
		curr_mural.Session.GameWon = false
	} else {
		curr_mural.Session.SubmittedAnswer = curr_mural.Session.SelectedAnswer
		if service.GetCorrectAnswer(curr_mural.Game.Answers).Movie.ID != curr_mural.Session.SelectedAnswer.Movie.ID {
			curr_mural.Session.CurrentScore = 0
			curr_mural.Session.GameWon = false
		}
	}

	// computer before we do stuff to this game
	game_shareable := service.ComputeShareable(curr_mural.Session, curr_mural.Game, curr_mural.UserData)

	var tiles [][]model.Tile
	var flipped int
	for _, row := range curr_mural.Session.Board.Tiles {
		var tile_row []model.Tile
		for _, tile := range row {
			if tile.Flipped {
				flipped += 1
			}

			tile := model.Tile{
				Penalty:  tile.Penalty,
				Selected: false,
				I:        tile.I,
				J:        tile.J,
				Flipped:  true,
			}

			tile_row = append(tile_row, tile)
		}

		tiles = append(tiles, tile_row)
	}

	curr_mural.Session.Board.Tiles = tiles
	curr_mural.Session.SessionStats = model.SessionStats{
		Score: curr_mural.Session.CurrentScore,
	}

	curr_mural.Session.SessionStatus = model.SESSION_OVER
	curr_mural.Session.SessionStats.Shareable = game_shareable

	// now that we have stats, let's add them to the database
	// db.DAL.SetStatsForUser(user_key, curr_mural.Session.SessionStats, curr_mural.Game)
	// db.DAL.SetGameSessionForUser(curr_mural.Session)

	// do analytics stuff
	// api.AnalyticsController.RegisterEvent(api.EVENT_SUBMIT, c.Request())
	return c.Render(http.StatusOK, "game-board.html", curr_mural)
}
