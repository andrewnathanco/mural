package worker

import (
	"log/slog"
	"math/rand"
	"mural/app"
	"mural/config"
	"mural/db"
	"time"
)

type MuralWorker struct {
	MuralService app.MuralService
}

func (mw MuralWorker) SetupNewGame() {
	slog.Info("RESETTING GAME SESSIONS")
	err := mw.MuralService.DAL.DeleteSessions()
	if err != nil {
		slog.Error(err.Error())
	}

	if mw.MuralService.Config.Env == config.EnvTest {
		mw.MuralService.Config.TodayTheme = config.ThemeRandom
	} else {
		mw.MuralService.Config.TodayTheme = config.GetTodayThemeDefault()
	}

	last_game, err := mw.MuralService.DAL.GetCurrentGame(mw.MuralService.Config)
	if err != nil {
		slog.Error(err.Error())
	}

	// start generating options
	_, err = mw.MuralService.DAL.SetNewCorrectOption(mw.MuralService.Config)
	if err != nil {
		slog.Error(err.Error())
	}

	_, err = mw.MuralService.DAL.SetNewEasyModeOptions(mw.MuralService.Config)
	if err != nil {
		slog.Error(err.Error())
	}

	// end the last game
	last_game.GameStatus = db.GAME_OVER
	mw.MuralService.DAL.UpsertGame(last_game)

	// start building the new one
	new_game := db.Game{
		GameKey:     last_game.GameKey + 1,
		OptionOrder: rand.Intn(4),
		PlayedOn:    time.Now(),
		Theme:       mw.MuralService.Config.TodayTheme,
		GameStatus:  db.GAME_CURRENT,
	}

	mw.MuralService.DAL.UpsertGame(new_game)
}
