package model

import (
	"time"

	"github.com/google/uuid"
)

type Movie struct {
	Name string
	ID uuid.UUID
	Poster string
	Description string
	ReleaseDate time.Time
}

type Tile struct {
	Penalty int
	Selected bool
	Flipped bool
	I int
	J int
}

type Board struct {
	Size int
	Tiles [][]Tile
}

type Answer struct {
	Movie
	IsCorrect bool
	Selected bool
}

type Game struct {
	CurrentScore int
	Selected *Tile
	SelectedTile *Answer
	Board Board
	Current Movie
	ScoreZones int
	Answers []Answer
}

