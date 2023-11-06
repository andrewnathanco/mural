package sql

import (
	"database/sql"
	"log/slog"
	"mural/db"
	"mural/model"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDAL struct {
	DB *sql.DB
}

func NewSQLiteDal(file string) (*SQLiteDAL, error) {
	err := createFileIfNotExists(file)
	if err != nil {
		slog.Error(err.Error())
		return nil, db.ErrCreateDatabaseFile
	}

	database, err := sql.Open("sqlite3", file)
	if err != nil {
		slog.Error(err.Error())
		return nil, db.ErrConnectToDatabase
	}

	err = database.Ping()
	if err != nil {
		return nil, db.ErrPingDatabase
	}

	// setup schema
	err = setupMuralSchema(database)
	if err != nil {
		return nil, db.ErrSetupGameSchema
	}

	return &SQLiteDAL{
		DB: database,
	 }, nil
}


func (dal *SQLiteDAL) GetGameSessionForUser(
	user_key string,
) (*model.Session, error) {
	session, err := getSession(user_key, dal)

	if err != nil  { 
		new_session, err := newSessionForUser(user_key, dal)
		if err != nil {
			return nil, err
		}

		err = insertSession(user_key, new_session, dal)
		if err != nil {
			return nil, err
		}

		session = new_session
	}

	return session, nil
}

func (dal *SQLiteDAL) SetGameSessionForUser(game_session model.Session) error {
	return insertSession(game_session.UserKey, &game_session, dal)
}

func (dal *SQLiteDAL) ResetGameSessions() error {
	return resetSessions(dal)
}

func (dal *SQLiteDAL) CacheAnswersInDatabase(answers []model.Answer) (error) {
	return cacheAnswers(answers, dal)
}

func (dal *SQLiteDAL) GetCurrentGameInfo() (*model.Game, error) {
	return getCurrentGameInfo(dal)
}

func (dal *SQLiteDAL) RedlistAnswer(answer model.Answer) error {
	// 
	current_game, err := dal.GetCurrentGameInfo()
	if err != nil {
		return err
	}

	return redlistAnswer(answer, *current_game, dal)
}

func (dal *SQLiteDAL) GetCurrentMoviePageFromDB() (*int, error) {
	return getCurrentMoviePageFromDB(dal)
}

func (dal *SQLiteDAL) SetCurrentMoviePageFromDB() (error) {
	current_movie_page, err := dal.GetCurrentMoviePageFromDB()
	if err != nil {
		return err
	}

	new_movie_page := *current_movie_page + 1
	return setCurrentMoviePageFromDB(new_movie_page, dal)
}

func (dal *SQLiteDAL) GetRandomAnswers() ([]model.Answer, error) {
	return getRandomAnswers(dal)
}

func (dal *SQLiteDAL) SetNewCurrentGame(current_game model.Game) (error) {
	return setNewCurrentGame(current_game, dal)
}

func (dal *SQLiteDAL) SetupMetadata() (error) {
	return setupMetadata(dal)
}
