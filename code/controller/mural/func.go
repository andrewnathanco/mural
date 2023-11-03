package mural

import (
	"mural/controller/shared"
	"mural/model"
	"time"
)


func getVersion() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// functions
func mod(a, b int) int {
    return a % b
}

type FlipButton struct {
	Button shared.Button
	SelectedTile *model.Tile
}

func newFlipButton(
	text string,
	disabled bool,
	tile *model.Tile,
) FlipButton {
	return FlipButton{
		Button: shared.Button{
			Text: text,
			Disabled: disabled,
		},
		SelectedTile: tile,
	}
}

type SelectItem struct {
	Answer model.Answer
	Disabled bool
}

func newSelectItem(answer model.Answer, disabled bool) SelectItem {
	return SelectItem{
		Answer: answer,
		Disabled: disabled,
	}
}

type SelectTile struct {
	Tile model.Tile
	Disabled bool
}

func newSelectTile(tile model.Tile, disabled bool) SelectTile {
	return SelectTile{
		Tile: tile,
		Disabled: disabled,
	}
}