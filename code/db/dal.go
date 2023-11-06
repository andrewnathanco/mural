package db

import (
	"fmt"
	"mural/model"
)

var (
	DAL IDAL

	// errors
	ErrConnectToDatabase = fmt.Errorf("could not connect to database")
	ErrCreateDatabaseFile = fmt.Errorf("could not create database file")
	ErrPingDatabase = fmt.Errorf("could not ping database")
	ErrSetupGameSchema = fmt.Errorf("could not setup game schema")
	ErrGettingBoardFromDB = fmt.Errorf("could not get board from db")
	ErrBoardNotFound = fmt.Errorf("no board set")
	ErrCastingBoard = fmt.Errorf("board not in correct format")
	ErrSettingCurrentGames = fmt.Errorf("could not set all boards")
)

type IDAL interface {
	GetGameSessionForUser(string) (*model.Session, error)
	SetGameSessionForUser(model.Session)  error
	ResetGameSessions()  error

	// metadata
	SetupMetadata()  error

	// answer stuff
	CacheAnswersInDatabase([]model.Answer) (error)
	RedlistAnswer(model.Answer) error
	GetCurrentMoviePageFromDB() (*int, error)
	SetCurrentMoviePageFromDB() (error)
	GetRandomAnswers() ([]model.Answer, error)

	GetCurrentGameInfo() (*model.Game, error)
	SetNewCurrentGame(model.Game) (error)
} 