package service

import (
	"log/slog"
	"math/rand"
	"mural/model"
	"os"
	"strings"
	"time"

	_ "time/tzdata"
)

func ResetSelected(all_tiles [][]model.Tile) [][]model.Tile {
	first_row := all_tiles[0]
	size := len(first_row)
	new_tiles := model.NewTiles(size)

	for _, row := range all_tiles {
		for _, tile := range row {
			tile := model.Tile{
				Penalty:  tile.Penalty,
				I:        tile.I,
				J:        tile.J,
				Selected: false,
				Flipped:  tile.Flipped,
			}

			new_tiles[tile.I][tile.J] = tile
		}
	}

	return new_tiles
}

func GetCorrectAnswer(answers []model.Answer) model.Answer {
	var answer model.Answer
	for _, a := range answers {
		if a.IsCorrect {
			answer = a
		}
	}

	return answer
}

func GetCurrentMural(
	user_key string,
) (*model.Mural, error) {
	// current_game, err := db.DAL.GetCurrentGameInfo()
	// if err != nil {
	// 	return nil, fmt.Errorf("could not get current game: %w", err)
	// }

	// current_session, err := db.DAL.GetGameSessionForUser(user_key)
	// if err != nil {
	// 	return nil, fmt.Errorf("could not get current session: %w", err)
	// }

	// user_data, err := db.DAL.GetUserData(user_key)
	// if err != nil {
	// 	return nil, fmt.Errorf("could not get user data: %w", err)
	// }

	// number_of_sessions, err := db.DAL.GetNumberOfSessions()
	// if err != nil {
	// 	return nil, fmt.Errorf("could not get number of session: %w", err)
	// }

	// current_game.NumberOfSessions = number_of_sessions
	// user_stats, _ := db.DAL.GetStatsForUser(user_key)
	// return &model.Mural{
	// 	Game: *current_game,
	// 	Session: *current_session,
	// 	UserStats: user_stats,
	// 	UserData: *user_data,
	// }, nil
	return nil, nil
}

func GetCurrentDecade() string {
	if strings.EqualFold(os.Getenv("ENV"), "dev") {
		decades := []string{
			"2020s",
			"2010s",
			"2000s",
			"1990s",
			"1980s",
			"1970s",
		}

		return decades[rand.Intn(len(decades))]
	}

	current_day := time.Now().Weekday()
	loc, _ := time.LoadLocation("America/New_York")

	if loc != nil {
		slog.Info(loc.String())
		current_day = time.Now().In(loc).Weekday()
	}

	switch current_day {
	case time.Monday:
		return "2020s"
	case time.Tuesday:
		return "2010s"
	case time.Wednesday:
		return "2000s"
	case time.Thursday:
		return "1990s"
	case time.Friday:
		return "1980s"
	case time.Saturday:
		return "1970s"
	default:
		// Sunday or any other day
		return "Random"
	}
}
