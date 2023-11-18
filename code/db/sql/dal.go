package sql

import (
	"database/sql"
	"fmt"
	"log/slog"
	"math/rand"
	"mural/config"
	"mural/db"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/samber/lo"
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

	_, err = dal.DB.Exec(createUsersTable)
	if err != nil {
		return err
	}

	_, err = dal.DB.Exec(createGameStatsTable)
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
				GameKey:     1,
				OptionOrder: rand.Intn(4),
				PlayedOn:    time.Now(),
				GameStatus:  db.GAME_CURRENT,
				Theme:       mur_conf.TodayTheme,
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
	err := dal.DB.Get(&number_of_sessions, getNumberOfSessionsPlayed, db.SESSION_WON, db.SESSION_LOST)
	return number_of_sessions, err
}

func (dal *SQLiteDAL) PopulateTiles(
	size int,
) error {
	tiles := generateGrid(size)
	_, err := dal.DB.NamedExec(insertTilesQuery, tiles)
	return err
}

func (dal *SQLiteDAL) SelectTileForUser(session_tile db.SessionTile) error {
	_, err := dal.DB.Exec(updateOtherSelectedTiles, db.TILE_DEFAULT, db.TILE_SELECTED)
	if err != nil {
		return err
	}
	return dal.SaveTileStatusForUser(session_tile)
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

func (dal *SQLiteDAL) GetScoreForUser(
	mural_conf config.MuralConfig,
	user_key string,
) (int, error) {
	board, err := dal.GetBoardForUser(mural_conf, user_key)
	if err != nil {
		return mural_conf.MaxScore, err
	}

	score := mural_conf.MaxScore
	for _, row := range board {
		for _, tile := range row {
			if tile.SessionTileStatus == db.TILE_FLIPPED {
				score -= tile.Penalty
			}
		}
	}

	return score, nil
}

func (dal *SQLiteDAL) SaveMovies(movies []db.Movie) error {
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

	correct_option, err := dal.GetCorrectOption()
	if err != nil {
		return mural, err
	}

	easy_options, err := dal.GetEasyModeOptions()
	if err != nil {
		return mural, err
	}

	options := insertOptionAtIndex(easy_options, correct_option, game.OptionOrder)
	game.CorrectOption = correct_option
	game.EasyModeOptions = options

	session, err := dal.GetSessionForUser(user_key)
	if err != nil {
		return mural, err
	}

	option, err := dal.GetOptionByKey(session.OptionKey)
	if err != sql.ErrNoRows {
		if err != nil {
			return mural, err
		}

		session.Option = option
	}

	board, err := dal.GetBoardForUser(mur_conf, user_key)
	if err != nil {
		return mural, err
	}

	score, err := dal.GetScoreForUser(mur_conf, user_key)
	if err != nil {
		return mural, err
	}

	session.Board = board
	session.CurrentScore = &score
	number_of_sessions, err := dal.GetNumberOfSessionsPlayed()
	if err != nil {
		return mural, err
	}

	// now get the user
	user, err := dal.GetUserByUserKey(user_key)
	if err != nil {
		return mural, err
	}

	user_stats, err := dal.GetWeeklyStatsForUser(user_key)
	if err != nil {
		return mural, err
	}

	current_streak, max_streak, err := dal.GetStreaksForUser(user_key)
	if err != nil {
		return mural, err
	}

	games_played, err := dal.GetTotalGamesPlayedByUser(user_key)
	if err != nil {
		return mural, err
	}

	user.WeeklyStats = user_stats
	user.CurrentStreak = current_streak
	user.MaxStreak = max_streak
	user.GamesPlayed = games_played

	// get back the game
	mural.Game = game
	mural.Session = session
	mural.User = user
	mural.Version = mur_conf.Version
	mural.NumberOfSessionsPlayed = number_of_sessions

	mural.BoardState = db.BOARD_NORMAL
	return mural, nil
}

func insertOptionAtIndex(options []db.Option, option db.Option, index int) []db.Option {
	if index < 0 || index > len(options) {
		return options
	}

	var new_options []db.Option
	new_options = append(new_options, options[:index]...)
	new_options = append(new_options, option)
	new_options = append(new_options, options[index:]...)

	return new_options
}

func (dal *SQLiteDAL) DeleteSessions() error {
	_, err := dal.DB.Exec(deleteSessions)
	if err != nil {
		return err
	}
	_, err = dal.DB.Exec(deleteSessionTiles)
	return err
}

func (dal *SQLiteDAL) UpsertOption(
	option db.Option,
) (int64, error) {
	// upsert option
	res, err := dal.DB.NamedExec(upsertOption, option)
	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

// TODO: add proper unit testing for this
func (dal *SQLiteDAL) GetRandomAvailableMovies(
	mur_conf config.MuralConfig,
	number int,
) ([]db.Movie, error) {
	movies := []db.Movie{}
	switch mur_conf.TodayTheme {
	case config.Theme2020:
		err := dal.DB.Select(
			&movies,
			getRandomMovie,
			775,
			getSQLDecade(mur_conf.TodayTheme),
			number,
		)
		if err != nil {
			return movies, err
		}
	case config.Theme2010:
		err := dal.DB.Select(
			&movies,
			getRandomMovie,
			1000,
			getSQLDecade(mur_conf.TodayTheme),
			number,
		)
		if err != nil {
			return movies, err
		}
	case config.Theme2000:
		err := dal.DB.Select(
			&movies,
			getRandomMovie,
			1000,
			getSQLDecade(mur_conf.TodayTheme),
			number,
		)
		if err != nil {
			return movies, err
		}
	case config.Theme1990:
		err := dal.DB.Select(
			&movies,
			getRandomMovie,
			450,
			getSQLDecade(mur_conf.TodayTheme),
			number,
		)
		if err != nil {
			return movies, err
		}
	case config.Theme1980:
		err := dal.DB.Select(
			&movies,
			getRandomMovie,
			320,
			getSQLDecade(mur_conf.TodayTheme),
			number,
		)
		if err != nil {
			return movies, err
		}
	case config.Theme1970:
		err := dal.DB.Select(
			&movies,
			getRandomMovie,
			250,
			getSQLDecade(mur_conf.TodayTheme),
			number,
		)
		if err != nil {
			return movies, err
		}
	default:
		err := dal.DB.Select(
			&movies,
			getRandomMovie,
			1500,
			getSQLDecade(config.ThemeRandom),
			number,
		)
		if err != nil {
			return movies, err
		}
	}

	// upsert option
	return movies, nil
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

	// update old movies
	_, err = dal.DB.Exec(resetOptionByStatus, db.OPTION_USED_CORRECT, db.OPTION_CORRECT)
	if err != sql.ErrNoRows {
		if err != nil {
			return option, err
		}
	}

	movie, err := dal.GetRandomAvailableMovies(mur_conf, 1)
	if err != nil {
		return option, err
	}

	if len(movie) != 1 {
		return option, fmt.Errorf("did not get 1 random answer")
	}

	new_option := db.Option{
		GameKey:      game.GameKey,
		OptionStatus: db.OPTION_CORRECT,
		Movie:        movie[0],
	}

	id, err := dal.UpsertOption(new_option)
	new_option.OptionKey = id
	return new_option, err
}

func (dal *SQLiteDAL) SetNewEasyModeOptions(mur_conf config.MuralConfig) ([]db.Option, error) {
	options := []db.Option{}
	// get current game
	game, err := dal.GetCurrentGame(mur_conf)
	if err != nil {
		return options, err
	}

	// update old movies
	_, err = dal.DB.Exec(resetOptionByStatus, db.OPTION_USED_EASY, db.OPTION_EASY_MODE)
	if err != nil {
		return options, err
	}

	// get random movie
	movies, err := dal.GetRandomAvailableMovies(mur_conf, 3)
	if err != nil {
		return options, err
	}

	if len(movies) != 3 {
		return options, fmt.Errorf("did not get 3 random answers")
	}

	new_options := []db.Option{}
	for _, movie := range movies {
		option := db.Option{
			OptionStatus: db.OPTION_EASY_MODE,
			GameKey:      game.GameKey,
			Movie:        movie,
		}

		id, err := dal.UpsertOption(option)
		if err != nil {
			return options, err
		}
		option.OptionKey = id
		new_options = append(new_options, option)
	}

	return new_options, err
}

func (dal *SQLiteDAL) GetMovieByMovieKey(movie_key int) (db.Movie, error) {
	movie := db.Movie{}
	err := dal.DB.Get(&movie, getMovieBykey, movie_key)
	return movie, err
}

func (dal *SQLiteDAL) GetOptionByMovie(movie_key int) (db.Option, error) {
	option, err := dal.GetCorrectOption()
	if err != nil {
		return option, err
	}

	if option.MovieKey == movie_key {
		return option, nil
	}

	movie, err := dal.GetMovieByMovieKey(movie_key)
	if err != nil {
		return option, err
	}

	// filler for the frontend
	option = db.Option{
		OptionStatus: db.OPTION_EMPTY,
		Movie:        movie,
	}

	id, err := dal.UpsertOption(option)
	if err != nil {
		return option, err
	}

	// get random movie
	option.OptionKey = id
	return option, err
}

func (dal *SQLiteDAL) GetCorrectOption() (db.Option, error) {
	var option db.Option
	err := dal.DB.Get(&option, getOptionByStatus, db.OPTION_CORRECT)
	if err != sql.ErrNoRows {
		if err != nil {
			return option, err
		}
	}

	// get random movie
	return option, err
}

func (dal *SQLiteDAL) GetEasyModeOptions() ([]db.Option, error) {
	var options []db.Option
	err := dal.DB.Select(&options, getOptionByStatus, db.OPTION_EASY_MODE)
	if err != sql.ErrNoRows {
		if err != nil {
			return options, err
		}
	}

	if len(options) != 3 {
		return options, fmt.Errorf("need 3 options")
	}

	// get random movie
	return options, err
}

func (dal *SQLiteDAL) GetBoardForUser(mural_conf config.MuralConfig, user_key string) ([][]db.SessionTile, error) {
	board := make([][]db.SessionTile, mural_conf.BoardWidth)
	for row_num := range board {
		board[row_num] = make([]db.SessionTile, mural_conf.BoardWidth)
		for col_num := range board[row_num] {
			tile, err := dal.GetSessionTileForUser(row_num, col_num, user_key)
			if err != nil {
				return board, err
			}

			board[row_num][col_num] = tile
		}
	}

	return board, nil
}

func (dal *SQLiteDAL) UpsertUser(user db.User) error {
	_, err := dal.DB.NamedExec(upsertUser, user)
	return err
}

func (dal *SQLiteDAL) GetUserByUserKey(user_key string) (db.User, error) {
	user := db.User{}
	err := dal.DB.Get(&user, getUserByKey, user_key)
	if err == sql.ErrNoRows {
		// create a new session
		user.UserKey = user_key
		user.GameType = db.REGULAR_MODE
		err = dal.UpsertUser(user)
	}

	return user, err
}

func (dal *SQLiteDAL) GetOptionByKey(option_key int64) (db.Option, error) {
	option := db.Option{}
	err := dal.DB.Get(&option, getOptionByKey, option_key)
	return option, err
}

func (dal *SQLiteDAL) GetOptionsByQuery(query string) ([]db.Option, error) {
	options := []db.Option{}
	movies := []db.Movie{}
	err := dal.DB.Select(&movies, queryMovies, query)
	if err != nil {
		return options, err
	}

	for _, movie := range movies {
		option := db.Option{
			OptionStatus: db.OPTION_EMPTY,
			Movie:        movie,
		}

		options = append(options, option)
	}

	return options, nil
}

func (dal *SQLiteDAL) UpsertGameStat(stat db.GameStat) error {
	_, err := dal.DB.NamedExec(upsertGameStat, stat)
	return err
}

func (dal *SQLiteDAL) GetTotalGamesPlayedByUser(user_key string) (int, error) {
	var games_played int
	err := dal.DB.Get(&games_played, getTotalGamesPlayedByUser, user_key)
	return games_played, err
}

// Returns the current streak first and the max streak second.
func (dal *SQLiteDAL) GetStreaksForUser(user_key string) (int, int, error) {
	var game_stats []db.GameStat

	// TODO come back to this
	err := dal.DB.Select(&game_stats, getAllGamesStatsForUser, user_key)
	if err == sql.ErrNoRows || len(game_stats) == 0 {
		return 0, 0, nil
	}

	if err != nil {
		return 0, 0, err
	}

	has_current := game_stats[0].GameStatus == db.GAME_CURRENT
	current_streak := 0
	max_streak := 0
	for i := 0; i < len(game_stats); i++ {
		if i == len(game_stats)-1 {
			continue
		}

		this_game_key := game_stats[i].GameKey
		last_game_key := game_stats[i+1].GameKey
		if has_current && (this_game_key-1 == last_game_key) {
			current_streak += 1
		} else {
			has_current = false
		}

		if this_game_key-1 == last_game_key {
			max_streak += 1
		} else {
			max_streak = 0
		}
	}

	// we need to add one here because this just tracks the bit inbetween the games, not the games themselves
	return current_streak + 1, max_streak + 1, err
}

func isDateInThisWeek(targetDate time.Time) bool {
	// Get the current time
	now := time.Now()

	// Calculate the start of the week (Sunday)
	week_start := now.AddDate(0, 0, -int(now.Weekday()))

	// Calculate the end of the week (Saturday)
	week_end := week_start.AddDate(0, 0, 6)

	// Check if the target date is within the current week
	return targetDate.After(week_start) && targetDate.Before(week_end)
}

// Returns the current streak first and the max streak second.
func (dal *SQLiteDAL) GetWeeklyStatsForUser(user_key string) (map[string]map[string]db.DayStat, error) {
	var game_stats []db.GameStat
	day_stat := db.DayStat{
		BestScore:    nil,
		AverageScore: nil,
		WeeklyScore:  nil,
	}

	weekly_stats := map[string]map[string]db.DayStat{
		db.REGULAR_MODE: {},
		db.EASY_MODE:    {},
	}

	game_map_by_day_easy := map[string][]db.GameStat{}
	game_map_by_day_regular := map[string][]db.GameStat{}

	// TODO come back to this
	err := dal.DB.Select(&game_stats, getAllGamesStatsForUser, user_key)
	if err == sql.ErrNoRows {
		return weekly_stats, nil
	}

	if err != nil {
		return weekly_stats, err
	}

	for _, game := range game_stats {
		// only get the weekly score if we don't already have it, theses games are sorted so we can trust this will always be the last one
		if game.GameType == db.EASY_MODE {
			// lets build out date set
			current_games := lo.ValueOr(game_map_by_day_easy, game.PlayedOn.Weekday().String(), []db.GameStat{})
			current_games = append(current_games, game)
			game_map_by_day_easy[game.PlayedOn.Weekday().String()] = current_games

			_, ok := weekly_stats[db.EASY_MODE][game.PlayedOn.Weekday().String()]
			if isDateInThisWeek(game.PlayedOn) && !ok && game.GameKey != 0 {
				day_stat.WeeklyScore = game.Score
				weekly_stats[db.EASY_MODE][game.PlayedOn.Weekday().String()] = day_stat
			}
		}

		if game.GameType == db.REGULAR_MODE {
			// lets build out date set
			current_games := lo.ValueOr(game_map_by_day_regular, game.PlayedOn.Weekday().String(), []db.GameStat{})
			current_games = append(current_games, game)
			game_map_by_day_regular[game.PlayedOn.Weekday().String()] = current_games

			_, ok := weekly_stats[db.REGULAR_MODE][game.PlayedOn.Weekday().String()]
			if isDateInThisWeek(game.PlayedOn) && !ok && game.GameKey != 0 {
				day_stat.WeeklyScore = game.Score
				weekly_stats[db.REGULAR_MODE][game.PlayedOn.Weekday().String()] = day_stat
			}
		}
	}

	games_by_type := map[string]map[string][]db.GameStat{
		db.EASY_MODE:    game_map_by_day_easy,
		db.REGULAR_MODE: game_map_by_day_regular,
	}

	for type_name, game_type := range games_by_type {
		for day, game_stats := range game_type {
			day_stat = weekly_stats[type_name][day]

			// get the average score
			score_sum := lo.SumBy[db.GameStat](game_stats, func(item db.GameStat) int {
				return lo.FromPtr(item.Score)
			})

			score_avg := score_sum / len(game_stats)
			day_stat.AverageScore = &score_avg

			// get the best score
			score_best := lo.MaxBy[db.GameStat](game_stats, func(a, b db.GameStat) bool {
				// could probably be improved, this is kinda a hack
				return lo.FromPtrOr(a.Score, -1000) > lo.FromPtrOr(b.Score, -1000)
			})

			day_stat.BestScore = score_best.Score
			weekly_stats[type_name][day] = day_stat
		}
	}

	return weekly_stats, nil
}
