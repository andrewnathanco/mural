package routes

import (
	"mural/controller/mural/service"
	"mural/db"
	"mural/middleware"
	"mural/model"
	"net/http"

	"github.com/labstack/echo/v4"
)


func SubmitAnswer(c echo.Context) error {
	user_key := middleware.GetUserKeyFromContext(c)
	curr_mural, err := service.GetCurrentMural(user_key)
    if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
    }

	curr_mural.Session.SubmittedAnswer = curr_mural.Session.SelectedAnswer
	if service.GetCorrectAnswer(curr_mural.Game.Answers).ID != curr_mural.Session.SelectedAnswer.ID {
		curr_mural.Session.CurrentScore = 0
	}
	// computer before we do stuff to this game
	game_shareable := service.ComputeShareable(curr_mural.Session, curr_mural.Game) 

	var tiles [][]model.Tile
	var flipped int
	for _, row := range curr_mural.Session.Board.Tiles {
		var tile_row []model.Tile
		for _, tile := range row {
			if tile.Flipped {
				flipped += 1
			}

			tile := model.Tile{
				Penalty: tile.Penalty,
				Selected: false,
				I: tile.I,
				J: tile.J,
				Flipped: true,
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
	db.DAL.SetStatsForUser(user_key, curr_mural.Session.SessionStats, curr_mural.Game)
	db.DAL.SetGameSessionForUser(curr_mural.Session)
	return c.Render(http.StatusOK, "game-board.html", curr_mural)
}
