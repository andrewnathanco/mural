package db

import "time"

type Game struct {
	GameKey    int       `json:"game_key" db:"game_key"`
	Theme      string    `json:"theme" db:"theme"`
	PlayedOn   time.Time `json:"played_on" db:"played_on"`
	GameStatus string    `json:"game_status" db:"game_status"`
}

const (
	GAME_CURRENT = "GAME_CURRENT"
	GAME_OVER    = "GAME_OVER"
)

const (
	Theme2020 = "2020"
	Theme2010 = "2010"
	Theme2000 = "2000"
	Theme1990 = "1990"
	Theme1980 = "1980"
	Theme1970 = "1970"
	Theme1960 = "random"
)

var (
	ThemeOptions = []string{
		Theme2020,
		Theme2010,
		Theme2000,
		Theme1990,
		Theme1980,
		Theme1970,
		Theme1960,
	}
)

type Session struct {
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
	TileKey    int    `json:"tile_key" db:"tile_key"`
	SessionKey int    `json:"session_key" db:"session_key"`
	TileStatus string `json:"tile_status" db:"tile_status"`
}

const (
	TILE_INIT     = "TILE_INIT"
	TILE_SELECTED = "TILE_SELECTED"
	TILE_FLIPPED  = "TILE_FLIPPED"
)

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
	// this is an abstract reference so that if we start doing other option type we can
	ReferenceKey int    `json:"reference_key" db:"reference_key"`
	GameKey      int    `json:"game_key" db:"game_key"`
	OptionStatus string `json:"option_status" db:"option_status"`
}

const (
	OPTION_USED = "OPTION_USED"
	// this is tough to name, basically indicate if it was used as a choice for easy mode
	OPTION_EASY_MODE = "OPTION_EASY_MODE"
)
