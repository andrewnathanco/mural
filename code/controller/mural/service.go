package mural

import (
	"mural/model"
)

// this needs an even number to be populated
func populateTileZones(
	max int, 
	size int, 
	level int,
	all_tiles [][]model.Tile, 
) (int, int, [][]model.Tile) {
	// break case
	if size == 0 {
		return max, 0, all_tiles
	}

	// lets get all of the tiles
	i := 0
	for i < size {
		j := 0
		for j < size {
			// top
			all_tiles[i + level][level] = model.Tile{
				Penalty: level + 1,
				Selected: false,
				Flipped: false,
			};

			// left
			all_tiles[level][j + level] = model.Tile{
				Penalty: level + 1,
				Selected: false,
				Flipped: false,
			};

			// right
			all_tiles[max -  1 - level][j + level] = model.Tile{
				Penalty: level + 1,
				Selected: false,
				Flipped: false,
			};

			// bottom
			all_tiles[i + level][max - 1 - level] = model.Tile{
				Penalty: level + 1,
				Selected: false,
				Flipped: false,
			};

			j+= 1

		}
		i+= 1
	}

	return populateTileZones(max, size - 2, level + 1, all_tiles)
}

func newGameBoard(size int) (*model.Board) {
	// need to make tiles
	tiles := make([][]model.Tile, size)

	// need to make rows
	for i := range tiles {
		tiles[i] = make([]model.Tile, size)
	}

	_, _, tiles =  populateTileZones(size, size, 0, tiles)

	return &model.Board{
		Size: size,
		Tiles: tiles,
	}
}

