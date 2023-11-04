package service

import (
	"mural/model"
)

func ResetSelected(all_tiles [][]model.Tile) [][]model.Tile {
	first_row := all_tiles[0]
	size := len(first_row)
	new_tiles := newTiles(size)

	for _, row := range all_tiles {
		for _, tile := range row {
			tile := model.Tile{
				Penalty: tile.Penalty,
				I: tile.I,
				J: tile.J,
				Selected:  false,
				Flipped: tile.Flipped,
			}

			new_tiles[tile.I][tile.J] = tile
		}
	}

	return new_tiles
}




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
			// left
			penalty := (level + 1)  * 3
			all_tiles[i + level][level] = model.Tile{
				I: i + level,
				J: level,
				Penalty: penalty,
				Selected: false,
				Flipped: false,
			};

			// top
			all_tiles[level][j + level] = model.Tile{
				I: level,
				J: j + level,
				Penalty: penalty,
				Selected: false,
				Flipped: false,
			};

			// bottom
			all_tiles[max -  1 - level][j + level] = model.Tile{
				I: max -  1 - level,
				J: j + level,
				Penalty: penalty,
				Selected: false,
				Flipped: false,
			};

			// right
			all_tiles[i + level][max - 1 - level] = model.Tile{
				I: i + level,
				J: max -  1 - level,
				Penalty: penalty,
				Selected: false,
				Flipped: false,
			};

			j+= 1

		}
		i+= 1
	}

	return populateTileZones(max, size - 2, level + 1, all_tiles)
}

func newTiles(size int) ([][]model.Tile) {
	// need to make tiles
	tiles := make([][]model.Tile, size)

	// need to make rows
	for i := range tiles {
		tiles[i] = make([]model.Tile, size)
	}

	_, _, tiles =  populateTileZones(size, size, 0, tiles)

	return tiles
}

func NewGameBoard(size int) (*model.Board) {
	return &model.Board{
		Size: size,
		Tiles: newTiles(size),
	}
}
