package sql

import (
	"math/rand"
	"mural/config"
	"mural/db"
	"mural/db/sql/test"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

var (
	DAL    *SQLiteDAL
	Config config.MuralConfig
)

func init() {
	Config = config.MuralConfig{
		BoardWidth:       10,
		TodayTheme:       config.Theme1970,
		DatabaseFile:     "./test/mural_test.db",
		MigrationsFolder: "./migrations",
	}

	// setup database
	var err error
	DAL, err = NewSQLiteDal(Config)
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

	current_game, err := DAL.GetOrCreateNewGame(Config)

	assert.NoError(t, err)
	assert.Equal(t, current_game.GameKey, game.GameKey)
	t.Cleanup(func() {
		DAL.DB.MustExec("delete from games")
	})
}

func TestCurrentGameNone(t *testing.T) {
	current_game, err := DAL.GetOrCreateNewGame(Config)
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

	current_game, err := DAL.GetOrCreateNewGame(Config)

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
			SessionStatus:     db.SESSION_LOST,
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

func TestUpsertOption(t *testing.T) {
	assert.NoError(t, DAL.SaveMovies(test.AllMovies))

	movie := db.Movie{}
	assert.NoError(t, DAL.DB.Get(&movie, "select * from movies where id = ?", test.MovBlueBeetle.ID))

	option := db.Option{
		OptionStatus: db.OPTION_CORRECT,
		Movie:        movie,
	}

	_, err := DAL.UpsertOption(option)
	assert.NoError(t, err)
	found_option := db.Option{}
	assert.NoError(t, DAL.DB.Get(&found_option, "select * from options where movie_key = ?", movie.MovieKey))
	assert.Equal(t, option.MovieKey, found_option.MovieKey)

	// creat the tile
	t.Cleanup(func() {
		DAL.DB.MustExec("delete from movies")
		DAL.DB.MustExec("delete from options")
	})
}

// test random movie with used option
func TestGetRandomMovieUseOption(t *testing.T) {
	config := config.MuralConfig{
		TodayTheme: config.Theme2020,
	}

	// need to populate with 2 popular movies
	movies := []db.Movie{
		test.MovBlueBeetle,
		test.MovMissionImpossible,
	}

	movie := db.Movie{}
	assert.NoError(t, DAL.SaveMovies(movies))
	assert.NoError(t, DAL.DB.Get(&movie, "select * from movies where id = ?", test.MovBlueBeetle.ID))

	option := db.Option{
		OptionStatus: db.OPTION_USED_CORRECT,
		Movie:        movie,
	}
	_, err := DAL.UpsertOption(option)
	assert.NoError(t, err)
	rand_movies, err := DAL.GetRandomAvailableMovies(config, 1)
	assert.NoError(t, err)
	assert.NotEqual(t, movie.MovieKey, rand_movies[0].MovieKey)
}

// test random movie with used option
func TestSaveCorrectOption(t *testing.T) {
	assert.NoError(t, DAL.SaveMovies(test.AllMovies))

	option, err := DAL.SetNewCorrectOption(config.MuralConfig{
		TodayTheme: config.Theme2020,
	},
		nil,
	)

	assert.NoError(t, err)
	assert.Equal(t, db.OPTION_CORRECT, option.OptionStatus)
	assert.NotNil(t, option.MovieKey)

	found_option := db.Option{}
	DAL.DB.Get(&found_option, `
		select * 
		from options 
		inner join movies 
		on movies.movie_key = options.movie_key
		where option_key = ?
	`, option.OptionKey)
	assert.Equal(t, option.OptionKey, found_option.OptionKey)
	assert.NotNil(t, option.MovieKey)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from movies")
		DAL.DB.MustExec("delete from options")
	})
}

// test random movie with used option
func TestGetEasyModeOptions(t *testing.T) {
	assert.NoError(t, DAL.SaveMovies(test.AllMovies))

	options, err := DAL.SetNewEasyModeOptions(config.MuralConfig{
		TodayTheme: config.Theme2020,
	})

	assert.NoError(t, err)
	assert.Len(t, options, 3)

	found_options := []db.Option{}
	assert.NoError(t, DAL.DB.Select(&found_options, `
		select * 
		from options 
		inner join movies 
		on movies.movie_key = options.movie_key
		where option_status = ?
	`, db.OPTION_EASY_MODE))
	assert.Len(t, found_options, 3)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from movies")
		DAL.DB.MustExec("delete from options")
	})
}

func TestUpsertUser(t *testing.T) {
	user_key := uuid.New().String()

	user := db.User{
		UserKey:  user_key,
		GameType: db.REGULAR_MODE,
	}

	assert.NoError(t, DAL.UpsertUser(user))

	var found_user db.User
	assert.NoError(t, DAL.DB.Get(&found_user, "select * from users where user_key = ?", user_key))
	assert.Equal(t, user_key, found_user.UserKey)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from users")
	})
}

func TestGetUserByID(t *testing.T) {
	user_key := uuid.New().String()
	user := db.User{
		UserKey:  user_key,
		GameType: db.REGULAR_MODE,
	}

	assert.NoError(t, DAL.UpsertUser(user))
	user, err := DAL.GetUserByUserKey(user_key)

	assert.NoError(t, err)
	assert.Equal(t, user_key, user.UserKey)
	assert.NotEmpty(t, user.UserKey)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from users")
	})
}

// test game stats
func TestUpsertGameStat(t *testing.T) {
	user_key := uuid.New().String()
	user := db.User{
		UserKey:  user_key,
		GameType: db.REGULAR_MODE,
	}

	assert.NoError(t, DAL.UpsertUser(user))

	game_key := 1
	game := db.Game{
		GameKey:    game_key,
		GameStatus: db.GAME_CURRENT,
		PlayedOn:   time.Now(),
		Theme:      config.Theme1980,
	}

	assert.NoError(t, DAL.UpsertGame(game))

	stat := db.GameStat{
		Game:          game,
		SessionStatus: db.SESSION_LOST,
		Score:         nil,
		UserKey:       user_key,
	}

	assert.NoError(t, DAL.UpsertGameStat(stat))
	var found_stat db.GameStat
	assert.NoError(t, DAL.DB.Get(&found_stat, "select * from game_stats where user_key = ? and game_key = ?", user_key, game_key))
	assert.Equal(t, user_key, found_stat.UserKey)
	assert.Equal(t, game_key, found_stat.GameKey)
	assert.Empty(t, found_stat.Score)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from users")
		DAL.DB.MustExec("delete from games")
		DAL.DB.MustExec("delete from game_stats")
	})
}

func TestGetTotalGamesPlayedByUser(t *testing.T) {
	user_key := uuid.New().String()
	user := db.User{
		UserKey:  user_key,
		GameType: db.REGULAR_MODE,
	}

	assert.NoError(t, DAL.UpsertUser(user))

	total_games := 5
	for i := 1; i <= total_games; i++ {
		game_key := i
		game := db.Game{
			GameKey:    game_key,
			GameStatus: db.GAME_CURRENT,
			PlayedOn:   time.Now(),
			Theme:      config.Theme1980,
		}

		assert.NoError(t, DAL.UpsertGame(game))

		score := i
		stat := db.GameStat{
			Game:          game,
			SessionStatus: db.SESSION_LOST,
			Score:         &score,
			UserKey:       user_key,
		}

		assert.NoError(t, DAL.UpsertGameStat(stat))
	}

	total, err := DAL.GetTotalGamesPlayedByUser(user_key)
	assert.NoError(t, err)
	assert.Equal(t, total_games, total)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from users")
		DAL.DB.MustExec("delete from games")
		DAL.DB.MustExec("delete from game_stats")
	})
}

func TestGetStreaks(t *testing.T) {
	user_key := uuid.New().String()
	user := db.User{
		UserKey:  user_key,
		GameType: db.REGULAR_MODE,
	}

	assert.NoError(t, DAL.UpsertUser(user))

	total_games := 10
	for i := 1; i <= total_games; i++ {
		game_status := db.GAME_OVER
		if i == total_games {
			game_status = db.GAME_CURRENT
		}

		game_key := i
		game := db.Game{
			GameKey:    game_key,
			GameStatus: game_status,
			PlayedOn:   time.Now(),
			Theme:      config.Theme1980,
		}

		assert.NoError(t, DAL.UpsertGame(game))

		score := i
		if i == total_games || i == 9 || i == 8 ||
			i == 5 || i == 4 || i == 3 || i == 2 || i == 1 {
			stat := db.GameStat{
				Game:          game,
				SessionStatus: db.SESSION_LOST,
				Score:         &score,
				UserKey:       user_key,
			}

			assert.NoError(t, DAL.UpsertGameStat(stat))
		}
	}

	current_streak, max_streak, err := DAL.GetStreaksForUser(user_key)
	assert.NoError(t, err)
	assert.Equal(t, 3, current_streak)
	assert.Equal(t, 5, max_streak)

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from users")
		DAL.DB.MustExec("delete from games")
		DAL.DB.MustExec("delete from game_stats")
	})
}

func TestGetWeeklyStats(t *testing.T) {
	// create new user
	user_key := uuid.New().String()
	user := db.User{
		UserKey:  user_key,
		GameType: db.REGULAR_MODE,
	}

	assert.NoError(t, DAL.UpsertUser(user))

	// create some games and stats
	total_games := 10
	scores := []int{}
	for i := 0; i < total_games; i++ {
		game_status := db.GAME_OVER
		if i == total_games {
			game_status = db.GAME_CURRENT
		}

		game_key := i + 1
		date := time.Now().AddDate(0, 0, i*-7)
		game := db.Game{
			GameKey:    game_key,
			GameStatus: game_status,
			PlayedOn:   date,
			Theme:      config.Theme1980,
		}

		assert.NoError(t, DAL.UpsertGame(game))

		score := rand.Intn(32)
		stat := db.GameStat{
			Game:          game,
			SessionStatus: db.SESSION_LOST,
			Score:         &score,
			GameType:      db.REGULAR_MODE,
			UserKey:       user_key,
		}

		assert.NoError(t, DAL.UpsertGameStat(stat))
		scores = append(scores, score)
	}

	sum := lo.Reduce[int](scores, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)

	avg := sum / len(scores)
	best := lo.Max(scores)

	stats, err := DAL.GetWeeklyStatsForUser(user_key)
	assert.NoError(t, err)
	assert.Equal(t, scores[0], lo.FromPtr(stats[db.REGULAR_MODE][time.Now().Weekday().String()].WeeklyScore))

	assert.Equal(t, avg, lo.FromPtr(stats[db.REGULAR_MODE][time.Now().Weekday().String()].AverageScore))
	assert.Equal(t, best, lo.FromPtr(stats[db.REGULAR_MODE][time.Now().Weekday().String()].BestScore))

	t.Cleanup(func() {
		DAL.DB.MustExec("delete from users")
		DAL.DB.MustExec("delete from games")
		DAL.DB.MustExec("delete from game_stats")
	})
}
