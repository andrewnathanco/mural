package db

import (
	"time"

	"github.com/ryanbradynd05/go-tmdb"
)

type MuralMeta struct {
	SystemKey         int `json:"system_key" db:"system_key"`
	LastTMDBMoviePage int `json:"last_tmdb_movie_page" db:"last_tmdb_movie_page"`
}

func NewMuralMeta(last_tmdb_movie_page int) MuralMeta {
	return MuralMeta{
		SystemKey:         1,
		LastTMDBMoviePage: last_tmdb_movie_page,
	}
}

type Game struct {
	GameKey         int       `json:"game_key" db:"game_key"`
	OptionOrder     int       `json:"option_order" db:"option_order"`
	Theme           string    `json:"theme" db:"theme"`
	PlayedOn        time.Time `json:"played_on" db:"played_on"`
	GameStatus      string    `json:"game_status" db:"game_status"`
	CorrectOption   Option
	EasyModeOptions []Option
}

const (
	GAME_CURRENT = "GAME_CURRENT"
	GAME_OVER    = "GAME_OVER"
)

type Session struct {
	CurrentScore int
	Board        [][]SessionTile
	SessionKey   int    `json:"session_key" db:"session_key"`
	UserKey      string `json:"user_key" db:"user_key"`
	Option
	SelectedOptionKey int    `json:"selected_option_key" db:"selected_option_key"`
	SessionStatus     string `json:"session_status" db:"session_status"`
}

const (
	SESSION_INIT    = "SESSION_INIT"
	SESSION_STARTED = "SESSION_STARTED"
	SESSION_LOST    = "SESSION_LOST"
	SESSION_WON     = "SESSION_WON"
)

type Tile struct {
	TileKey   int `json:"tile_key" db:"tile_key"`
	RowNumber int `json:"row_number" db:"row_number"`
	ColNumber int `json:"col_number" db:"col_number"`
	Penalty   int `json:"penalty" db:"penalty"`
}

type SessionTile struct {
	Tile
	SessionKey        int    `json:"session_key" db:"session_key"`
	SessionTileStatus string `json:"tile_status" db:"tile_status"`
}

const (
	TILE_DEFAULT  = "TILE_DEFAULT"
	TILE_SELECTED = "TILE_SELECTED"
	TILE_FLIPPED  = "TILE_FLIPPED"
)

func ConvertShortToMovies(tmdbShort tmdb.MovieShort) Movie {
	return Movie{
		ID:            tmdbShort.ID,
		Title:         tmdbShort.Title,
		OriginalTitle: tmdbShort.OriginalTitle,
		ReleaseDate:   tmdbShort.ReleaseDate,
		Overview:      tmdbShort.Overview,
		VoteAverage:   tmdbShort.VoteAverage,
		VoteCount:     tmdbShort.VoteCount,
		Popularity:    tmdbShort.Popularity,
		Adult:         tmdbShort.Adult,
		Video:         tmdbShort.Video,
		BackdropPath:  tmdbShort.BackdropPath,
		PosterPath:    tmdbShort.PosterPath,
	}
}

type Movie struct {
	MovieKey      int     `db:"movie_key" json:"movie_key"`
	ID            int     `db:"id" json:"id"`
	Title         string  `db:"title" json:"title"`
	OriginalTitle string  `db:"original_title" json:"original_title"`
	ReleaseDate   string  `db:"release_date" json:"release_date"`
	Overview      string  `db:"overview" json:"overview"`
	VoteAverage   float32 `db:"vote_average" json:"vote_average"`
	VoteCount     uint32  `db:"vote_count" json:"vote_count"`
	Popularity    float32 `db:"popularity" json:"popularity"`
	Adult         bool    `db:"adult" json:"adult"`
	Video         bool    `db:"video" json:"video"`
	BackdropPath  string  `db:"backdrop_path" json:"backdrop_path"`
	PosterPath    string  `db:"poster_path" json:"poster_path"`
}

type Option struct {
	OptionKey int64 `json:"option_key" db:"option_key"`
	Movie
	GameKey      int    `json:"game_key" db:"game_key"`
	OptionStatus string `json:"option_status" db:"option_status"`
}

type UserOption struct {
	OptionKey int `json:"option_key" db:"option_key"`
	Movie
	GameKey      int    `json:"game_key" db:"game_key"`
	OptionStatus string `json:"option_status" db:"option_status"`
}

const (
	OPTION_USED_CORRECT = "OPTION_USED_CORRECT"
	OPTION_CORRECT      = "OPTION_CORRECT"
	// this is tough to name, basically indicate if it was used as a choice for easy mode
	OPTION_EASY_MODE = "OPTION_EASY_MODE"
	OPTION_USED_EASY = "OPTION_USED_EASY"
	OPTION_EMPTY     = "OPTION_EMPTY"
)

type User struct {
	UserKey  string `json:"user_key" db:"user_key"`
	GameType string `json:"game_type" db:"game_type"`
	UserStats
}

type UserStats struct {
	WeeklyStats   map[string]DailyStat `json:"user_stats" `
	MaxStreak     int
	CurrentStreak int
	GamesPlayed   int
}

type DailyStat struct {
	Day     string `json:"day" db:"day"`
	UserKey string `json:"user_key" db:"user_key"`
	Best    *int   `json:"best" db:"best"`
	Week    *int   `json:"week" db:"week"`
}

const (
	DayMon = "Mon"
	DayTue = "Tue"
	DayWed = "Wed"
	DayThu = "Thu"
	DayFri = "Fri"
	DaySat = "Sat"
	DaySun = "Sun"
)

const (
	STAT_WEEK = "STAT_WEEK"
	STAT_BEST = "STAT_BEST"
)
const (
	REGULAR_MODE = "REGULAR_MODE"
	EASY_MODE    = "EASY_MODE"
)

// this is the package that everyone gets
type Mural struct {
	User                   User
	Session                Session
	Game                   Game
	Version                string
	NumberOfSessionsPlayed int
}
