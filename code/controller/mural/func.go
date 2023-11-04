package mural

import (
	"mural/controller/shared"
	"mural/model"
	"os"
)


func getVersion() string {
	return os.Getenv("VERSION")
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

type ShareButton struct {
	Button shared.Button
	Game model.Game
}

func newShareButton(
	text string,
	disabled bool,
	game model.Game,
) ShareButton {
	return ShareButton{
		Button: shared.Button{
			Text: text,
			Disabled: disabled,
		},
		Game: game,
	}
}

type StatsButton struct {
	Button shared.Button
	Game model.Game
}

func newStatsButton(
	text string,
	game model.Game,
) StatsButton {
	return StatsButton{
		Button: shared.Button{
			Text: text,
			Disabled: false,
		},
		Game: game,
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