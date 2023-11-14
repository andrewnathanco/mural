package sql

import (
	"database/sql"
	"log/slog"
	"mural/config"
	"mural/db"
	"os"
	"time"

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
	_, err := dal.DB.Exec(createMetaTable)
	if err != nil {
		return err
	}

	_, err = dal.DB.Exec(createGameTable)
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

	_, err = dal.DB.Exec(createMovieTable)
	if err != nil {
		return err
	}

	_, err = dal.DB.Exec(createOptionTable)
	if err != nil {
		return err
	}

	return nil
}

func (dal *SQLiteDAL) PingDatabase() error {
	return dal.DB.Ping()
}

func (dal *SQLiteDAL) UpsertMeta(meta db.MuralMeta) error {
	_, err := dal.DB.NamedExec(upsertMeta, meta)
	return err
}

func (dal *SQLiteDAL) GetMeta() (db.MuralMeta, error) {
	var meta db.MuralMeta
	err := dal.DB.Get(&meta, getMeta)
	if err == sql.ErrNoRows {

		meta = db.NewMuralMeta(1)
		err = dal.UpsertMeta(meta)
	}

	return meta, err
}

func (dal *SQLiteDAL) UpsertGame(game db.Game) error {
	_, err := dal.DB.NamedExec(upsertGameQuery, game)
	return err
}

func (dal *SQLiteDAL) GetCurrentGame(
	mur_conf config.MuralConfig,
) (db.Game, error) {
	game := db.Game{}
	err := dal.DB.Get(&game, getGameByStatus, db.GAME_CURRENT)
	if err == sql.ErrNoRows {
		err = dal.DB.Get(&game, getLastGame)
		if err == sql.ErrNoRows {
			game = db.Game{
				GameKey:    1,
				PlayedOn:   time.Now(),
				GameStatus: db.GAME_CURRENT,
				Theme:      mur_conf.TodayTheme,
			}

			err = dal.UpsertGame(game)
			return game, err
		}

		if err != nil {
			return game, err
		}

		game.Theme = mur_conf.TodayTheme
		game.GameKey = game.GameKey + 1
		game.GameStatus = db.GAME_CURRENT
		err = dal.UpsertGame(game)
	}

	return game, err
}

func (dal *SQLiteDAL) UpsertSession(session db.Session) error {
	_, err := dal.DB.NamedExec(upsertSession, session)
	return err
}

func (dal *SQLiteDAL) GetSessionForUser(user_key string) (db.Session, error) {
	session := db.Session{}
	err := dal.DB.Get(&session, getSessionByUser, user_key)
	if err == sql.ErrNoRows {
		// create a new session
		session.UserKey = user_key
		session.SessionStatus = db.SESSION_INIT
		err = dal.UpsertSession(session)
	}

	return session, err
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

func (dal *SQLiteDAL) SaveTileStatusForUser(session_tile db.SessionTile) error {
	_, err := dal.DB.NamedExec(upsertSessionTiles, session_tile)
	return err
}

func (dal *SQLiteDAL) GetTile(row int, col int) (db.Tile, error) {
	tile := db.Tile{}
	err := dal.DB.Get(&tile, getTile, row, col)
	return tile, err
}

func (dal *SQLiteDAL) GetSessionTileForUser(row int, col int, user_key string) (db.SessionTile, error) {
	// step 1: try to get the selected tile
	session_tile := db.SessionTile{}
	err_session := dal.DB.Get(&session_tile, getSessionTileForUser, row, col, user_key)
	if err_session == sql.ErrNoRows {
		tile, err := dal.GetTile(row, col)
		if err != nil {
			return session_tile, err
		}

		user_sess, err := dal.GetSessionForUser(user_key)
		if err != nil {
			return session_tile, err
		}

		// step 2
		session_tile = db.SessionTile{
			SessionKey:        user_sess.SessionKey,
			Tile:              tile,
			SessionTileStatus: db.TILE_DEFAULT,
		}
	}

	return session_tile, nil
}

func (dal *SQLiteDAL) SaveMovies(movies []db.Movie) error {
	_, err := dal.DB.NamedExec(upsertMovie, movies)
	return err
}

func (dal *SQLiteDAL) GetOptionsByDecade(movies []db.Movie) error {
	_, err := dal.DB.NamedExec(upsertMovie, movies)
	return err
}

func (dal *SQLiteDAL) GetMuralForUser(
	user_key string,
	mur_conf config.MuralConfig,
) (db.Mural, error) {
	mural := db.Mural{}
	game, err := dal.GetCurrentGame(mur_conf)
	if err != nil {
		return mural, err
	}

	session, err := dal.GetSessionForUser(user_key)
	if err != nil {
		return mural, err
	}

	number_of_sessions, err := dal.GetNumberOfSessionsPlayed()
	if err != nil {
		return mural, err
	}

	// get back the game
	mural.Game = game
	mural.Session = session
	mural.Version = mur_conf.Version
	mural.NumberOfSessionsPlayed = number_of_sessions
	return mural, nil
}

func (dal *SQLiteDAL) ResetGame(
	mural_conf config.MuralConfig,
) {
	// delete all of the user sessions
}

func (dal *SQLiteDAL) DeleteSessions() error {
	_, err := dal.DB.Exec(deleteSessions)
	return err
}

func (dal *SQLiteDAL) UpsertOption(
	option db.Option,
) error {
	// upsert option
	_, err := dal.DB.NamedExec(upsertOption, option)
	return err
}

// TODO: add proper unit testing for this
func (dal *SQLiteDAL) GetRandomAvailableMovie(
	mur_conf config.MuralConfig,
) (db.Movie, error) {
	movie := db.Movie{}
	switch mur_conf.TodayTheme {
	case config.Theme2020:
		err := dal.DB.Get(
			&movie,
			getRandomMovie,
			775,
			getSQLDecade(mur_conf.TodayTheme),
		)
		if err != nil {
			return movie, err
		}
	case config.Theme2010:
		err := dal.DB.Get(
			&movie,
			getRandomMovie,
			1000,
			getSQLDecade(mur_conf.TodayTheme),
		)
		if err != nil {
			return movie, err
		}
	case config.Theme2000:
		err := dal.DB.Get(
			&movie,
			getRandomMovie,
			1000,
			getSQLDecade(mur_conf.TodayTheme),
		)
		if err != nil {
			return movie, err
		}
	case config.Theme1990:
		err := dal.DB.Get(
			&movie,
			getRandomMovie,
			450,
			getSQLDecade(mur_conf.TodayTheme),
		)
		if err != nil {
			return movie, err
		}
	case config.Theme1980:
		err := dal.DB.Get(
			&movie,
			getRandomMovie,
			320,
			getSQLDecade(mur_conf.TodayTheme),
		)
		if err != nil {
			return movie, err
		}
	case config.Theme1970:
		err := dal.DB.Get(
			&movie,
			getRandomMovie,
			250,
			getSQLDecade(mur_conf.TodayTheme),
		)
		if err != nil {
			return movie, err
		}

	}

	// upsert option
	return movie, nil
}

func getSQLDecade(current_decade string) string {
	decade_sql := ""
	if current_decade == config.ThemeRandom {
		decade_sql += "%"
	} else {
		decade_sql += replaceLastCharacter(current_decade, '%')
	}

	return decade_sql
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

func (dal *SQLiteDAL) SetNewCorrectOption(
	mur_conf config.MuralConfig,
) (db.Option, error) {
	option := db.Option{}
	// get current game
	game, err := dal.GetCurrentGame(mur_conf)
	if err != nil {
		return option, err
	}

	err = dal.DB.Get(&option, getCurrentCorrectOption, db.OPTION_CORRECT)
	if err == sql.ErrNoRows {
		// get random movie
		movie, err := dal.GetRandomAvailableMovie(mur_conf)
		if err != nil {
			return option, err
		}

		option = db.Option{
			OptionStatus: db.OPTION_CORRECT,
			GameKey:      game.GameKey,
			Movie:        movie,
		}

		err = dal.UpsertOption(option)
		return option, err
	}

	// reset the old one
	option.OptionStatus = db.OPTION_USED
	movie, err := dal.GetRandomAvailableMovie(mur_conf)
	if err != nil {
		return option, err
	}

	option = db.Option{
		GameKey:      game.GameKey,
		OptionStatus: db.OPTION_CORRECT,
		Movie:        movie,
	}

	return option, nil
}

func (dal *SQLiteDAL) SetNewEasyModeOptions(config config.MuralConfig) ([]db.Option, error) {
	options := []db.Option{}
	return options, nil
}
