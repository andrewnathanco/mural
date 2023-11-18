package mural

import (
	"fmt"
	"html/template"
	"mural/config"
	"mural/controller/shared"
	"mural/db"
	"mural/model"
	"strings"
	"time"

	"github.com/samber/lo"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func getNumberOfFlippedTiples(board model.Board) int {
	// we should be able to trust this, not just put an empty string
	number_of_flipped := 0
	for _, row := range board.Tiles {
		for _, tile := range row {
			if tile.Flipped {
				number_of_flipped += 1
			}

		}
	}

	return number_of_flipped
}

func getReleaseYear(movie db.Movie) string {
	// we should be able to trust this, not just put an empty string
	layout := "2006-01-02"
	release_date, err := time.Parse(layout, movie.ReleaseDate)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%d", release_date.Year())
}

func convertStringToHTML(text string) template.HTML {
	return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
}

func addCommaToNumber(number int) string {
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", number)
}

func getDate() string {
	return time.Now().Format(time.RFC3339)
}

// functions
func mod(a, b int) int {
	return a % b
}

// functions
func sub(a, b int) int {
	return a - b
}

// functions
func div(a, b int) int {
	return a % b
}

// functions
func bang(a bool) bool {
	return !a
}

type FlipButton struct {
	Button      shared.Button
	SessionTile db.SessionTile
}

func newFlipButton(
	text string,
	disabled bool,
	tile db.SessionTile,
) FlipButton {
	return FlipButton{
		Button: shared.Button{
			Text:     text,
			Disabled: disabled,
		},
		SessionTile: tile,
	}
}

type InfoButton struct {
	Button  shared.Button
	Session db.Session
}

func newInfoButton(
	text string,
	session db.Session,
) InfoButton {
	return InfoButton{
		Button: shared.Button{
			Text:     text,
			Disabled: false,
		},
		Session: session,
	}
}

type ShareButton struct {
	Button shared.Button
	Mural  db.Mural
}

func newShareButton(
	text string,
	disabled bool,
	mural db.Mural,
) ShareButton {
	return ShareButton{
		Button: shared.Button{
			Text:     text,
			Disabled: disabled,
		},
		Mural: mural,
	}
}

type StatsButton struct {
	Button  shared.Button
	Session db.Session
}

func newStatsButton(
	text string,
	session db.Session,
) StatsButton {
	return StatsButton{
		Button: shared.Button{
			Text:     text,
			Disabled: false,
		},
		Session: session,
	}
}

type SelectItem struct {
	Option   db.Option
	Disabled bool
}

func newSelectItem(option db.Option, disabled bool) SelectItem {
	return SelectItem{
		Option:   option,
		Disabled: disabled,
	}
}

type SelectTile struct {
	SessionTile db.SessionTile
	Disabled    bool
}

func newSelectTile(tile db.SessionTile, disabled bool) SelectTile {
	return SelectTile{
		SessionTile: tile,
		Disabled:    disabled,
	}
}

func getDecadeString(theme string) string {
	if theme != config.ThemeRandom {
		return theme + "s"
	}

	return theme
}

func getHaveString(sessions int) string {
	if sessions == 1 {
		return "Has"
	} else {
		return "Have"
	}
}

func getSelectedTileFromBoard(board [][]db.SessionTile) db.SessionTile {
	var selected_tile db.SessionTile
	for _, row := range board {
		for _, tile := range row {
			if tile.SessionTileStatus == db.TILE_SELECTED {
				selected_tile = tile
			}
		}
	}

	return selected_tile
}

func getShareable(
	mural db.Mural,
) string {
	header := "Mural"
	if mural.User.GameType == db.REGULAR_MODE {
		header += "*"
	}

	var score string
	if mural.Session.SessionStatus == db.SESSION_WON {
		score = fmt.Sprintf("%d", lo.FromPtr(mural.Session.CurrentScore))
	} else {
		score = "❎"
	}

	text := fmt.Sprintf("%s #%d Score: %s\n\n", header, mural.Game.GameKey, score)

	// need to make tiles
	for _, row := range mural.Session.Board {
		for _, tile := range row {
			if mural.User.GameType == db.REGULAR_MODE {
				if tile.SessionTileStatus == db.TILE_FLIPPED {
					text += "⬜"
				} else {
					text += "🟪"
				}
			} else {
				if tile.SessionTileStatus == db.TILE_FLIPPED {
					text += "⬜"
				} else {
					text += "🟩"
				}
			}
		}
		text += "\n"
	}

	text += "\nPlay at: mural.andrewnathan.net"
	return text
}
