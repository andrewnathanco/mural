package model

import (
	"time"

	"github.com/google/uuid"
)

type GameState string

const (
	GAME_INIT = "GAME_INIT"
	GAME_STARTED = "GAME_STARTED"
	GAME_OVER = "GAME_OVER"
)

type Movie struct {
	Name string `json:",omitempty"`
	ID uuid.UUID `json:",omitempty"`
	Poster string `json:",omitempty"`
	Description string `json:",omitempty"`
	ReleaseDate time.Time `json:",omitempty"`
}

type Tile struct {
	Penalty int 
	Selected bool
	Flipped bool 
	I int 
	J int 
}
type Board struct {
	Size int `json:",omitempty"`
	Tiles [][]Tile `json:",omitempty"`
}

type Answer struct {
	Movie `json:",omitempty"`
	IsCorrect bool
	Selected bool
}

type Game struct {
	GameKey string 

	// current metadata
	CurrentScore int
	Board Board 

	// answers
	Answers []Answer 
	TodayAnswer Movie 

	// current state
	SelectedTile *Tile `json:",omitempty"`
	SelectedAnswer *Answer `json:",omitempty"`
	SubmittedAnswer *Answer `json:",omitempty"`
	GameState GameState 
}

