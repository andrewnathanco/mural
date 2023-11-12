package sql

import (
	"mural/config"
	"mural/db"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	DAL *SQLiteDAL
)

func init() {
	// setup sessions
	os.Setenv("DATABASE_FILE", "./test/mural_test.db")

	// setup database
	var err error
	DAL, err = NewSQLiteDal(os.Getenv(config.EnvDatabasFile))
	config.Must(err)
}

// game stuff
func TestUpsertGame(t *testing.T) {
	game_key := 1
	game := db.Game{
		GameKey:    game_key,
		GameStatus: db.GAME_CURRENT,
		PlayedOn:   time.Now(),
		Theme:      db.Theme1980,
	}

	assert.NoError(t, DAL.UpsertGame(game))

	found_game := db.Game{}

	assert.NoError(t, DAL.DB.Get(&found_game, "select * from games where game_key = ?", game_key))
	assert.Equal(t, found_game.GameKey, game_key)
	DAL.DB.MustExec("delete from games")
}

func TestCurrentGame(t *testing.T) {
	game_key := 1
	game := db.Game{
		GameKey:    game_key,
		GameStatus: db.GAME_CURRENT,
		PlayedOn:   time.Now(),
		Theme:      db.Theme1980,
	}

	game_key_2 := 2
	game_2 := db.Game{
		GameKey:    game_key_2,
		GameStatus: db.GAME_OVER,
		PlayedOn:   time.Now(),
		Theme:      db.Theme1980,
	}

	assert.NoError(t, DAL.UpsertGame(game))
	assert.NoError(t, DAL.UpsertGame(game_2))

	current_game, err := DAL.GetCurrentGame()

	assert.NoError(t, err)
	assert.Equal(t, current_game.GameKey, game.GameKey)
	DAL.DB.MustExec("delete from games")
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
	DAL.DB.MustExec("delete from sessions")
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
	current_session, err := DAL.GetSessionByUser(user_key)

	assert.NoError(t, err)
	assert.Equal(t, user_key, current_session.UserKey)
	assert.NotEmpty(t, current_session.SessionKey)
	DAL.DB.MustExec("delete from sessions")
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
	DAL.DB.MustExec("delete from sessions")
}

func TestPopulateTiles(t *testing.T) {
	size := 10
	assert.NoError(t, DAL.PopulateTiles(size))
	var num_of_tiles int
	assert.NoError(t, DAL.DB.Get(&num_of_tiles, "select count(*) from tiles"))
	assert.Equal(t, size*size, num_of_tiles)
	DAL.DB.MustExec("delete from tiles")
}
