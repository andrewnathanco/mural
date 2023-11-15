package worker

import (
	"log/slog"
	"mural/app"
	"mural/db"
	"time"
)

type MuralWorker struct{}

func NewMuralWorker() MuralWorker {
	return MuralWorker{}
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
	if err != nil {
		slog.Error(err.Error())
	}

	// start generating options
	_, err = mur_serv.DAL.SetNewCorrectOption(mur_serv.Config)
	if err != nil {
		slog.Error(err.Error())
	}

	_, err = mur_serv.DAL.SetNewEasyModeOptions(mur_serv.Config)
	if err != nil {
		slog.Error(err.Error())
	}

	// end the last game
	last_game.GameStatus = db.GAME_OVER
	mur_serv.DAL.UpsertGame(last_game)

	// start building the new one
	new_game := db.Game{
		GameKey:    last_game.GameKey + 1,
		PlayedOn:   time.Now(),
		Theme:      mur_serv.Config.TodayTheme,
		GameStatus: db.GAME_CURRENT,
	}

	mur_serv.DAL.UpsertGame(new_game)
}
