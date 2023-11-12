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
	dal := SQLiteDAL{DB: database}
	err = dal.InitDB()
	return &dal, err
}

func (dal *SQLiteDAL) InitDB() error {
	_, err := dal.DB.Exec(createGameTable)
	if err != nil {
		return err
	}

	_, err = dal.DB.Exec(createSessionTable)
	if err != nil {
		return err
	}

	_, err = dal.DB.Exec(createTilesTables)
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

func (dal *SQLiteDAL) UpsertSession(session db.Session) error {
	_, err := dal.DB.NamedExec(upsertSession, session)
	return err
}

func (dal *SQLiteDAL) GetSessionByUser(user_key string) (*db.Session, error) {
	session := db.Session{}
	err := dal.DB.Get(&session, getSessionByUser, user_key)
	return &session, err
}

func (dal *SQLiteDAL) GetNumberOfSessionsPlayed() (int, error) {
	var number_of_sessions int
	err := dal.DB.Get(&number_of_sessions, getNumberOfSessionsPlayed, db.SESSION_OVER)
	return number_of_sessions, err
}

func (dal *SQLiteDAL) PopulateTiles(
	size int,
) error {
	tiles := generateGrid(size)
	_, err := dal.DB.NamedExec(insertTilesQuery, tiles)
	return err
}

func generateGrid(size int) []db.Tile {
	if size <= 0 {
		return nil
	}

	var tiles []db.Tile
	tileKey := 1

	for ring := 0; ring < size; ring++ {
		penalty := -(ring * 3) // Calculate the penalty based on the ring
		for col := ring; col < size-ring; col++ {
			tiles = append(tiles, db.Tile{
				TileKey:   tileKey,
				RowNumber: ring,
				ColNumber: col,
				Penalty:   penalty,
			})
			tileKey++
		}
		for row := ring + 1; row < size-ring; row++ {
			tiles = append(tiles, db.Tile{
				TileKey:   tileKey,
				RowNumber: row,
				ColNumber: size - ring - 1,
				Penalty:   penalty,
			})
			tileKey++
		}
		for col := size - ring - 2; col >= ring; col-- {
			tiles = append(tiles, db.Tile{
				TileKey:   tileKey,
				RowNumber: size - ring - 1,
				ColNumber: col,
				Penalty:   penalty,
			})
			tileKey++
		}
		for row := size - ring - 2; row > ring; row-- {
			tiles = append(tiles, db.Tile{
				TileKey:   tileKey,
				RowNumber: row,
				ColNumber: ring,
				Penalty:   penalty,
			})
			tileKey++
		}
	}

	return tiles
}
