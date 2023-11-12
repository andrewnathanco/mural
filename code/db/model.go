package db

import "time"

type Game struct {
	GameKey int `json:"game_key" db:"game_key"`
	Theme string `json:"theme" db:"theme"`
	PlayedOn time.Time `json:"played_on" db:"played_on"`
	GameStatus string `json:"game_status" db:"game_status"`
}

const (
	GAME_CURRENT = "GAME_CURRENT"
	GAME_OVER = "GAME_OVER"
)

type Theme interface {
	GetOptions() ([]Option, error)
}

type Session struct {
	SessionKey int `json:"session_key" db:"session_key"`
	UserKey string `json:"user_key" db:"user_key"`
	SelectedOptionKey int `json:"selected_option_key" db:"selected_option_key"`
	SessionStatus int `json:"session_status" db:"session_status"`
}

const (
	SESSION_INIT = "SESSION_INIT"
	SESSION_SELECTED = "SESSION_SELECTED"
	SESSION_OVER = "SESSION_OVER"
)

type Tiles struct {
	TileKey int `json:"tile_key" db:"tile_key"`
	RowNumber int `json:"row_number" db:"row_number"`
	ColNumber int `json:"col_number" db:"col_number"`
	CellData interface{} `json:"cell_data"`
	Penalty int `json:"penalty" db:"penalty"`
}

type SessionTiles struct {
	TileKey int `json:"tile_key" db:"tile_key"`
	SessionKey int `json:"session_key" db:"session_key"`
	TileStatus string `json:"tile_status" db:"tile_status"`
}

const (
	TILE_INIT = "TILE_INIT"
	TILE_SELECTED = "TILE_SELECTED"
	TILE_FLIPPED = "TILE_FLIPPED"
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
	ReferenceKey int `json:"reference_key" db:"reference_key"`
	GameKey int `json:"game_key" db:"game_key"`
	OptionStatus string `json:"option_status" db:"option_status"`
}

const (
	OPTION_USED = "OPTION_USED"
	// this is tough to name, basically indicate if it was used as a choice for easy mode
	OPTION_EASY_MODE = "OPTION_EASY_MODE"
)

