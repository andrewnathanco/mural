package mural

import (
	"fmt"
	"html/template"
	"mural/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	board_size = 6
)
func mod(a, b int) int {
    return a % b
}

type Poster struct {
	Name string
	URL string
}

type Tile struct {
	Penalty int
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

// this needs an even number to be populated
func PopulateTileZones(
	max int, 
	size int, 
	level int,
	all_tiles [][]Tile, 
	err error,
) (int, int, [][]Tile, error) {
	if size % 2 != 0 {
		return 0, 0, nil, fmt.Errorf("needs an even number")
	}

	// break case
	if size == 0 {
		return max, 0, all_tiles, err
	}

	// lets get all of the tiles
	i := 0
	for i < size {
		j := 0
		for j < size {
			// top
			all_tiles[i + level][level] = Tile{
				Penalty: level + 1,
				Flipped: false,
			};

			// left
			all_tiles[level][j + level] = Tile{
				Penalty: level + 1,
				Flipped: false,
			};

			// right
			all_tiles[max -  1 - level][j + level] = Tile{
				Penalty: level + 1,
				Flipped: false,
			};

			// bottom
			all_tiles[i + level][max - 1 - level] = Tile{
				Penalty: level + 1,
				Flipped: false,
			};

			j+= 1

		}
		i+= 1
	}

	return PopulateTileZones(max, size - 2, level + 1, all_tiles, err)
}

func NewGameBoard(size int) (*Board, error) {
	// need to make tiles
	tiles := make([][]Tile, size)

	// need to make rows
	for i := range tiles {
        tiles[i] = make([]Tile, size)
    }

	_, _, tiles, err :=  PopulateTileZones(size, size, 0, tiles, nil)
	if err != nil {
		return nil, err
	}

	return &Board{
		Size: size,
		Tiles: tiles,
	}, nil
}

func NewMuralController() (*controller.TemplateController, func(c echo.Context) error, error) {
	mod_map := template.FuncMap{
        "mod": mod,
    }

	mural_template := template.Must(template.New("mural_template").Funcs(mod_map).ParseFiles("view/mural/mural.html", "view/mural/game/game-board.html"))
	new_template_controller := controller.TemplateController{
		Template: mural_template,
		Name: "mural.html",
	}
	
	return &new_template_controller, GetMural, nil
}

func GetMural(c echo.Context) error {
	board, err := NewGameBoard(board_size)
	if err != nil {
		return err
	}

	game := Game{
		CurrentScore: 56,
		Board: *board,
		Poster: Poster{
			Name: "Talk To Me",
			URL: "https://image.tmdb.org/t/p/w1280//kdPMUMJzyYAc4roD52qavX0nLIC.jpg",
		},
	}

	return c.Render(http.StatusOK, "mural.html", game)
}
