package worker

import (
	"log/slog"
	"math/rand"
	"mural/app"
	"mural/controller/mural/service"
	"mural/db"
	"mural/model"
	"time"
)

type MuralWorker struct{}

func NewMuralWorker() MuralWorker {
	return MuralWorker{}
}

func getSQLDecade() string {
	decade_str := service.GetCurrentDecade()
	currentDay := time.Now().Weekday()

	decade_sql := ""
	if currentDay == time.Sunday {
		decade_sql += "%"
	} else {
		decade := decade_str[0 : len(decade_str)-1]
		decade_sql += replaceLastCharacter(decade, '%')
	}

	return decade_sql
}

func (mw MuralWorker) SetupNewGame(
	mur_serv app.MuralService,
) {
	slog.Info("RESETTING GAME SESSIONS")
	err := mur_serv.DAL.DeleteSessions()
	if err != nil {
		slog.Error(err.Error())
	}

	// then lets game the game we just played
	last_game, err := mur_serv.DAL.GetCurrentGame(mur_serv.Config)

	// start generating options
	mur_serv.DAL.SetNewCorrectOption(mur_serv.Config)
	mur_serv.DAL.SetNewEasyModeOptions(mur_serv.Config)

	// end the last game
	last_game.GameStatus = db.GAME_OVER
	mur_serv.DAL.UpsertGame(last_game)

	// start building the new one
	last_game.GameKey += 1
	last_game.PlayedOn = time.Now()
	last_game.Theme = mur_serv.Config.TodayTheme
	last_game.GameStatus = db.GAME_CURRENT
	mur_serv.DAL.UpsertGame(last_game)
}

func randomizeAnswers(a []model.Answer) {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}

func replaceLastCharacter(
	inputString string,
	newChar rune,
) string {
	if len(inputString) == 0 {
		return inputString // Return the original string if it's empty
	}

	// Convert the string to a rune slice to work with individual characters
	strRunes := []rune(inputString)

	// Update the last character
	strRunes[len(strRunes)-1] = newChar

	// Convert the rune slice back to a string
	return string(strRunes)
}
