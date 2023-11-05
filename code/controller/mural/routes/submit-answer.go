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
	game_key, err := middleware.GetGameKeyFromContext(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, "could not get game key")
	}

	current_game, err := db.DAL.GetCurrentGame(game_key)
    if err != nil {
		return c.String(http.StatusInternalServerError, "could not get current game")
    }


	var selected_answer model.Answer
	for _, a := range current_game.Answers {
		if a.Selected  {
			selected_answer = a
		}
	}

	current_game.SubmittedAnswer = &selected_answer
	if current_game.TodayAnswer.ID != selected_answer.ID {
		current_game.CurrentScore = 0
	}
	// computer before we do stuff to this game
	game_shareable := service.ComputeShareable(*current_game) 

	var tiles [][]model.Tile
	var flipped int
	for _, row := range current_game.Board.Tiles {
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

	current_game.Board.Tiles = tiles
	current_game.GameStats = model.GameStats{
		Score: current_game.CurrentScore,
		TilesFlipped: flipped,
	}

	current_game.GameState = model.GAME_OVER
	current_game.GameStats.Shareable = game_shareable

	db.DAL.SetCurrentGame(*current_game)
	return c.Render(http.StatusOK, "game-board.html", current_game)
}
