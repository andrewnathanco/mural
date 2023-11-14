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
	SessionTiles      []SessionTile
	SessionKey        int    `json:"session_key" db:"session_key"`
	UserKey           string `json:"user_key" db:"user_key"`
	SelectedOptionKey int    `json:"selected_option_key" db:"selected_option_key"`
	SessionStatus     string `json:"session_status" db:"session_status"`
}

const (
	SESSION_INIT     = "SESSION_INIT"
	SESSION_SELECTED = "SESSION_SELECTED"
	SESSION_OVER     = "SESSION_OVER"
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
	TILE_DEFAULT  = "TILE_INIT"
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
	OptionKey int `json:"option_key" db:"option_key"`
	Movie
	GameKey      int    `json:"game_key" db:"game_key"`
	OptionStatus string `json:"option_status" db:"option_status"`
}

const (
	OPTION_USED = "OPTION_CORRECT"
	// this is tough to name, basically indicate if it was used as a choice for easy mode
	OPTION_EASY_MODE = "OPTION_EASY_MODE"
)

type User struct {
	UserKey int    `json:"user_key" db:"user_key"`
	Name    string `json:"name" db:"name"`
	// this is an abstract reference so that if we start doing other option type we can
	GameType      string    `json:"game_type" db:"game_type"`
	LastPlayed    time.Time `json:"last_played" db:"last_played"`
	TotalScore    int       `json:"total_score" db:"total_score"`
	BestScore     int       `json:"best_score" db:"best_score"`
	CurrentStreak int       `json:"current_streak" db:"current_streak"`
	BestStreak    int       `json:"best_streak" db:"best_streak"`
}

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
