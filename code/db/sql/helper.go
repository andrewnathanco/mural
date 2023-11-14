package sql

import "mural/db"

func generateGrid(size int) []db.Tile {
	if size <= 0 {
		return nil
	}

	var tiles []db.Tile

	for ring := 0; ring < size; ring++ {
		penalty := (ring + 1) * 3 // Calculate the penalty based on the ring
		for col := ring; col < size-ring; col++ {
			tiles = append(tiles, db.Tile{
				RowNumber: ring,
				ColNumber: col,
				Penalty:   penalty,
			})
		}
		for row := ring + 1; row < size-ring; row++ {
			tiles = append(tiles, db.Tile{
				RowNumber: row,
				ColNumber: size - ring - 1,
				Penalty:   penalty,
			})
		}
		for col := size - ring - 2; col >= ring; col-- {
			tiles = append(tiles, db.Tile{
				RowNumber: size - ring - 1,
				ColNumber: col,
				Penalty:   penalty,
			})
		}
		for row := size - ring - 2; row > ring; row-- {
			tiles = append(tiles, db.Tile{
				RowNumber: row,
				ColNumber: ring,
				Penalty:   penalty,
			})
		}
	}

	return tiles
}
