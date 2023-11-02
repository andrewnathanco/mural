package model

type Poster struct {
	Name string
	URL string
}

type Tile struct {
	Penalty int
	Selected bool
	Flipped bool
}

type Board struct {
	Size int
	Tiles [][]Tile
}

type Game struct {
	CurrentScore int
	Board Board
	Poster Poster
	ScoreZones int
}

