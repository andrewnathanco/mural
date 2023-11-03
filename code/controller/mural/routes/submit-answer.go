package routes

import (
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

	var tiles [][]model.Tile
	for _, row := range current_game.Board.Tiles {
		var tile_row []model.Tile
		for _, tile := range row {
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
	current_game.SubmittedAnswer = &selected_answer
	current_game.GameState = model.GAME_OVER

	db.DAL.SetCurrentGame(*current_game)
	return c.Render(http.StatusOK, "game-board.html", current_game)
}
