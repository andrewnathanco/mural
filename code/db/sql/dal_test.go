package sql

import (
	"mural/config"
	"mural/db"
	"mural/db/sql/test"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	DAL    *SQLiteDAL
	Config config.MuralConfig
)

func init() {
	Config = config.MuralConfig{
		BoardWidth:   10,
		TodayTheme:   config.Theme1970,
		DatabaseFile: "./test/mural_test.db",
	}

	// setup database
	var err error
	DAL, err = NewSQLiteDal(Config.DatabaseFile)
	config.Must(err)
}

func TestUpsertMeta(t *testing.T) {
	first_page := 1
	second_page := 2
	meta := db.NewMuralMeta(first_page)
	assert.NoError(t, DAL.UpsertMeta(meta))
	var found_meta db.MuralMeta
	err := DAL.DB.Get(&found_meta, "select * from mural_meta", nil)
	assert.NoError(t, err)
	assert.Equal(t, first_page, found_meta.LastTMDBMoviePage)

	// now change the page
	found_meta.LastTMDBMoviePage = second_page
	assert.NoError(t, DAL.UpsertMeta(found_meta))
	err = DAL.DB.Get(&found_meta, "select * from mural_meta", nil)
	assert.NoError(t, err)
	assert.Equal(t, second_page, found_meta.LastTMDBMoviePage)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from mural_meta")
	})
}

func TestGetMeta(t *testing.T) {
	page := 2
	meta := db.NewMuralMeta(page)
	assert.NoError(t, DAL.UpsertMeta(meta))

	found_meta, err := DAL.GetMeta()
	assert.NoError(t, err)
	assert.Equal(t, found_meta.LastTMDBMoviePage, 2)
	assert.Equal(t, found_meta.SystemKey, 1)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from mural_meta")
	})
}

func TestGetMetaNoExists(t *testing.T) {
	found_meta, err := DAL.GetMeta()
	assert.NoError(t, err)
	assert.Equal(t, found_meta.LastTMDBMoviePage, 1)
	assert.Equal(t, found_meta.SystemKey, 1)
	t.Cleanup(func() {
		DAL.DB.MustExec("delete from mural_meta")
	})
}

// game stuff
func TestUpsertGame(t *testing.T) {
	game_key := 1
	game := db.Game{
		GameKey:    game_key,
		GameStatus: db.GAME_CURRENT,
		PlayedOn:   time.Now(),
		Theme:      config.Theme1980,
	}

	assert.NoError(t, DAL.UpsertGame(game))

	found_game := db.Game{}

	assert.NoError(t, DAL.DB.Get(&found_game, "select * from games where game_key = ?", game_key))
	assert.Equal(t, found_game.GameKey, game_key)
	t.Cleanup(func() {
		DAL.DB.MustExec("delete from games")
	})
}

func TestCurrentGame(t *testing.T) {
	game_key := 1
	game := db.Game{
		GameKey:    game_key,
		GameStatus: db.GAME_CURRENT,
		PlayedOn:   time.Now(),
		Theme:      config.Theme1980,
	}

	game_key_2 := 2
	game_2 := db.Game{
		GameKey:    game_key_2,
		GameStatus: db.GAME_OVER,
		PlayedOn:   time.Now(),
		Theme:      config.Theme1980,
	}

	assert.NoError(t, DAL.UpsertGame(game))
	assert.NoError(t, DAL.UpsertGame(game_2))

	current_game, err := DAL.GetCurrentGame(Config)

	assert.NoError(t, err)
	assert.Equal(t, current_game.GameKey, game.GameKey)
	t.Cleanup(func() {
		DAL.DB.MustExec("delete from games")
	})
}

func TestCurrentGameNone(t *testing.T) {
	current_game, err := DAL.GetCurrentGame(Config)
	assert.NoError(t, err)
	assert.Equal(t, current_game.GameKey, 1)

	// manually check to see if it exists
	found_game := db.Game{}
	assert.NoError(t, DAL.DB.Get(&found_game, getGameByStatus, db.GAME_CURRENT))
	assert.Equal(t, db.GAME_CURRENT, found_game.GameStatus)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from games")
	})
}

func TestCurrentGameNoCurrent(t *testing.T) {
	game_key := 1
	game := db.Game{
		GameKey:    game_key,
		GameStatus: db.GAME_OVER,
		PlayedOn:   time.Now(),
		Theme:      config.Theme1980,
	}

	game_key_2 := 2
	game_2 := db.Game{
		GameKey:    game_key_2,
		GameStatus: db.GAME_OVER,
		PlayedOn:   time.Now(),
		Theme:      config.Theme1980,
	}

	assert.NoError(t, DAL.UpsertGame(game))
	assert.NoError(t, DAL.UpsertGame(game_2))

	current_game, err := DAL.GetCurrentGame(Config)

	assert.NoError(t, err)
	assert.Equal(t, game_2.GameKey+1, current_game.GameKey)

	// manually check to see if it exists
	found_game := db.Game{}
	assert.NoError(t, DAL.DB.Get(&found_game, getGameByStatus, db.GAME_CURRENT))
	assert.Equal(t, db.GAME_CURRENT, found_game.GameStatus)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from games")
	})
}

// session stuff
func TestUpsertSession(t *testing.T) {
	session_key := 1
	user_key := uuid.New().String()
	selected_option := 1
	session := db.Session{
		SessionKey:        session_key,
		UserKey:           user_key,
		SessionStatus:     db.SESSION_INIT,
		SelectedOptionKey: selected_option,
	}

	assert.NoError(t, DAL.UpsertSession(session))

	found_session := db.Session{}

	assert.NoError(t, DAL.DB.Get(&found_session, "select * from sessions where session_key = ?", session_key))
	assert.Equal(t, found_session.SessionKey, session_key)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from sessions")
	})
}

func TestGetSessionByUser(t *testing.T) {
	user_key := uuid.New().String()
	selected_option := 1
	session := db.Session{
		UserKey:           user_key,
		SessionStatus:     db.SESSION_INIT,
		SelectedOptionKey: selected_option,
	}

	assert.NoError(t, DAL.UpsertSession(session))
	current_session, err := DAL.GetSessionForUser(user_key)

	assert.NoError(t, err)
	assert.Equal(t, user_key, current_session.UserKey)
	assert.NotEmpty(t, current_session.SessionKey)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from sessions")
	})
}

func TestGetNumberOfSessionsPlayed(t *testing.T) {
	// create session that have been played
	number_of_sess := 15
	for i := 0; i < number_of_sess; i++ {
		user_key := uuid.New().String()
		selected_option := 1
		session := db.Session{
			UserKey:           user_key,
			SessionStatus:     db.SESSION_OVER,
			SelectedOptionKey: selected_option,
		}

		assert.NoError(t, DAL.UpsertSession(session))
	}

	// create session that hasn't been started
	user_key := uuid.New().String()
	selected_option := 1
	session := db.Session{
		UserKey:           user_key,
		SessionStatus:     db.SESSION_INIT,
		SelectedOptionKey: selected_option,
	}

	assert.NoError(t, DAL.UpsertSession(session))

	number_of_sessions_played, err := DAL.GetNumberOfSessionsPlayed()

	assert.NoError(t, err)
	assert.Equal(t, number_of_sess, number_of_sessions_played)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from sessions")
	})
}

func TestPopulateTiles(t *testing.T) {
	size := 10
	assert.NoError(t, DAL.PopulateTiles(size))
	var num_of_tiles int
	assert.NoError(t, DAL.DB.Get(&num_of_tiles, "select count(*) from tiles"))
	assert.Equal(t, size*size, num_of_tiles)
	DAL.DB.MustExec("delete from tiles")
}

func TestUpsertUserSessionTile(t *testing.T) {
	session_key := 1
	tile_key := 1
	size := 10

	assert.NoError(t, DAL.PopulateTiles(size))
	row_num := 0
	col_num := 0
	tile, err := DAL.GetTile(row_num, col_num)
	assert.NoError(t, err)

	session_tile := db.SessionTile{
		SessionKey:        session_key,
		Tile:              tile,
		SessionTileStatus: db.TILE_FLIPPED,
	}
	assert.NoError(t, DAL.SaveTileStatusForUser(session_tile))

	var found_session_tile db.SessionTile
	assert.NoError(t, DAL.DB.Get(&found_session_tile, "select * from session_tiles where session_key = ? and tile_key = ?", session_key, tile_key))
	assert.Equal(t, session_key, found_session_tile.SessionKey)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from session_tiles")
		DAL.DB.MustExec("delete from tiles")
	})
}

func TestPenalty(t *testing.T) {
	size := 10
	assert.NoError(t, DAL.PopulateTiles(size))

	var penalty int
	assert.NoError(t, DAL.DB.Get(&penalty, "select penalty from tiles where row_number = ? and col_number = ?", 0, 0))

	assert.Equal(t, penalty, 3)
	t.Cleanup(func() {
		DAL.DB.MustExec("delete from tiles")
	})
}

func TestGetTile(t *testing.T) {
	size := 10
	assert.NoError(t, DAL.PopulateTiles(size))
	row_num := 0
	col_num := 0
	tile, err := DAL.GetTile(row_num, col_num)
	assert.NoError(t, err)

	assert.Equal(t, tile.RowNumber, row_num)
	assert.Equal(t, tile.ColNumber, col_num)
	t.Cleanup(func() {
		DAL.DB.MustExec("delete from tiles")
	})
}

func TestSaveSessionTileForUser(t *testing.T) {
	starting_status := db.TILE_SELECTED
	ending_status := db.TILE_FLIPPED

	// create the tile
	size := 10
	assert.NoError(t, DAL.PopulateTiles(size))
	row_num := 0
	col_num := 0
	tile, err := DAL.GetTile(row_num, col_num)
	assert.NoError(t, err)

	user_key := uuid.New().String()

	// create the session
	selected_option := 1
	session := db.Session{
		UserKey:           user_key,
		SessionStatus:     db.SESSION_INIT,
		SelectedOptionKey: selected_option,
	}

	assert.NoError(t, DAL.UpsertSession(session))
	found_session, err := DAL.GetSessionForUser(user_key)
	assert.NoError(t, err)

	// create session tile
	session_tile := db.SessionTile{
		SessionKey:        found_session.SessionKey,
		Tile:              tile,
		SessionTileStatus: starting_status,
	}
	assert.NoError(t, DAL.SaveTileStatusForUser(session_tile))

	// now that we found the tile, check to make sure its right
	found_tile, err := DAL.GetSessionTileForUser(row_num, col_num, user_key)
	assert.NoError(t, err)
	assert.Equal(t, starting_status, found_tile.SessionTileStatus)

	// this time it should be changed
	found_tile.SessionTileStatus = ending_status
	assert.NoError(t, DAL.SaveTileStatusForUser(found_tile))
	found_tile, err = DAL.GetSessionTileForUser(row_num, col_num, user_key)
	assert.NoError(t, err)
	assert.Equal(t, ending_status, found_tile.SessionTileStatus)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from session_tiles")
		DAL.DB.MustExec("delete from tiles")
		DAL.DB.MustExec("delete from sessions")
	})
}

func TestGetSessionTileSessionExists(t *testing.T) {
	// create the tile
	size := 10
	assert.NoError(t, DAL.PopulateTiles(size))
	row_num := 0
	col_num := 0
	tile, err := DAL.GetTile(row_num, col_num)
	assert.NoError(t, err)

	var user_keys []string
	var session_keys []int
	for i := 0; i < 4; i++ {
		user_key := uuid.New().String()
		user_keys = append(user_keys, user_key)

		// create the session
		selected_option := 1
		session := db.Session{
			UserKey:           user_key,
			SessionStatus:     db.SESSION_INIT,
			SelectedOptionKey: selected_option,
		}

		assert.NoError(t, DAL.UpsertSession(session))
		found_session, err := DAL.GetSessionForUser(user_key)
		assert.NoError(t, err)

		// create session tile
		session_keys = append(session_keys, found_session.SessionKey)
		session_tile := db.SessionTile{
			SessionKey:        found_session.SessionKey,
			Tile:              tile,
			SessionTileStatus: db.TILE_FLIPPED,
		}
		assert.NoError(t, DAL.SaveTileStatusForUser(session_tile))
	}

	found_tile, err := DAL.GetSessionTileForUser(row_num, col_num, user_keys[0])
	assert.NoError(t, err)

	assert.Equal(t, tile.TileKey, found_tile.Tile.TileKey)
	assert.Equal(t, session_keys[0], found_tile.SessionKey)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from session_tiles")
		DAL.DB.MustExec("delete from tiles")
		DAL.DB.MustExec("delete from sessions")
	})
}

func TestGetSessionTileSessionNoExist(t *testing.T) {
	// creat the tile
	size := 10
	assert.NoError(t, DAL.PopulateTiles(size))
	row_num := 0
	col_num := 0
	tile, err := DAL.GetTile(row_num, col_num)
	assert.NoError(t, err)

	// create the session
	user_key := uuid.New().String()
	selected_option := 1
	session := db.Session{
		UserKey:           user_key,
		SessionStatus:     db.SESSION_INIT,
		SelectedOptionKey: selected_option,
	}

	assert.NoError(t, DAL.UpsertSession(session))
	found_tile, err := DAL.GetSessionTileForUser(row_num, col_num, session.UserKey)
	assert.NoError(t, err)

	assert.Equal(t, tile.TileKey, found_tile.Tile.TileKey)
	assert.Equal(t, db.TILE_DEFAULT, found_tile.SessionTileStatus)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from session_tiles")
		DAL.DB.MustExec("delete from tiles")
	})
}

func TestSaveMovie(t *testing.T) {
	assert.NoError(t, DAL.SaveMovies(test.AllMovies))

	var movie_id int
	assert.NoError(t, DAL.DB.Get(&movie_id, "select id from movies where id = ?", test.MovMissionImpossible.ID))
	assert.Equal(t, test.MovMissionImpossible.ID, movie_id)

	// creat the tile
	t.Cleanup(func() {
		DAL.DB.MustExec("delete from movies")
	})
}

func TestRandomAvailableMovie(t *testing.T) {
	assert.NoError(t, DAL.SaveMovies(test.AllMovies))

	// creat the tile
	t.Cleanup(func() {
		DAL.DB.MustExec("delete from movies")
		DAL.DB.MustExec("delete from options")
	})
}

func TestUpsertOption(t *testing.T) {
	assert.NoError(t, DAL.SaveMovies(test.AllMovies))

	// creat the tile
	t.Cleanup(func() {
		DAL.DB.MustExec("delete from movies")
		DAL.DB.MustExec("delete from options")
	})
}
