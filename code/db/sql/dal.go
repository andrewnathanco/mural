package sql

import (
	"log/slog"
	"mural/db"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type SQLiteDAL struct {
	DB *sqlx.DB
}

func createFileIfNotExists(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		// File does not exist, so create it
		_, err = os.Create(filename)
		return err
	}
	return nil
}

func NewSQLiteDal(database_str string) (*SQLiteDAL, error) {
	err := createFileIfNotExists(database_str)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	database, err := sqlx.Open("sqlite3", database_str)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}

	// setup
	dal := SQLiteDAL{ DB: database }
	err = dal.InitDB()
	return &dal, err
}

func (dal *SQLiteDAL) InitDB() error {
	_, err := dal.DB.Exec(createGameQuery)
	if err != nil {
		return err
	}
	return nil
}

func (dal *SQLiteDAL) PingDatabase() error {
	return dal.DB.Ping()
}

func (dal *SQLiteDAL) UpsertGame(game db.Game) error {
	_, err := dal.DB.NamedExec(upsertGameQuery, game)
	return err
}

func (dal *SQLiteDAL) GetCurrentGame() (*db.Game, error) {
	game := db.Game{}
	err := dal.DB.Get(&game, getGameByStatus, db.GAME_CURRENT)
	return &game, err
}