package db

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"mural/api"
	"mural/controller/mural/cons"
	"mural/controller/mural/service"
	"mural/model"
	"os"

	_ "github.com/mattn/go-sqlite3"
)


const createGameTables string = `
	create table if not exists games (
	game_key string not null primary key,
	game string not null
);`

const createStatsTables string = `
	create table if not exists stats (
	game_key string not null primary key,
	stats string not null
);`

const upsertGame string = `
    insert or replace into games (game_key, game)
    values (?, ?)
	
`
const selectGame string = `
    select game from games
    where game_key = ?
`

type SQLiteDAL struct {
	DB *sql.DB
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

func createNewGame(
	game_key string,
) (*model.Game, error) {
	board := service.NewGameBoard(cons.BoardSize)
	correct_movie, answers, err := api.MovieController.GetAnswers()
	if err != nil {
		return nil, err
	}

	current_game := model.Game{
		GameKey: game_key,
		CurrentScore: cons.InititalScore,
		Board: *board,
		TodayAnswer: *correct_movie,
		Answers: answers,
		GameState: model.GAME_INIT,
	}

	return &current_game, nil
}

func insertGame(
	game_key string,
	game *model.Game,
	dal *SQLiteDAL,
) (error) {
	current_games_marshalled, err := json.Marshal(game)
	if err != nil {
		return err
	}


	_, err = dal.DB.Exec(upsertGame, game_key, string(current_games_marshalled))
	if err != nil {
		return err
	}

	return nil
}

func getGame(
	game_key string,
	dal *SQLiteDAL,
) (*model.Game, error) {
	var game_str string
	var game model.Game
	row := dal.DB.QueryRow(selectGame, game_key)
	err := row.Scan(&game_str)
	if err != nil  {
		return nil, err
	}

	err = json.Unmarshal([]byte(game_str), &game)
	if err != nil {
		return nil, err
	}

	return &game, nil
}


func setupGameSchema(
	db *sql.DB,
) error {
	_, err := db.Exec(createGameTables)
	if err != nil {
		return err
	}

	_, err = db.Exec(createStatsTables)
	if err != nil {
		return err
	}

	return nil
}

func NewSQLiteDal(file string) (*SQLiteDAL, error) {
	err := createFileIfNotExists(file)
	if err != nil {
		slog.Error(err.Error())
		return nil, ErrCreateDatabaseFile
	}

	db, err := sql.Open("sqlite3", file)
	if err != nil {
		slog.Error(err.Error())
		return nil, ErrConnectToDatabase
	}

	err = db.Ping()
	if err != nil {
		return nil, ErrPingDatabase
	}

	// setup schema
	err = setupGameSchema(db)
	if err != nil {
		return nil, ErrSetupGameSchema
	}

	return &SQLiteDAL{
		DB: db,
	 }, nil
}

func (dal *SQLiteDAL) InitGames () error {
	return nil
}


func (dal *SQLiteDAL) GetCurrentGames (
)(map[string]model.Game, error) {

	return nil, nil
}

func (dal *SQLiteDAL) GetCurrentGame(
	game_key string,
) (*model.Game, error) {
	
	game, err := getGame(game_key, dal)

	if err != nil  { 
		new_game, err := createNewGame(game_key)
		if err != nil {
			return nil, err
		}

		err = insertGame(game_key, new_game, dal)
		if err != nil {
			return nil, err
		}

		game = new_game
	}

	return game, nil
}

func (dal *SQLiteDAL) SetCurrentGame(current_game model.Game) error {
	return insertGame(current_game.GameKey, &current_game, dal)
}