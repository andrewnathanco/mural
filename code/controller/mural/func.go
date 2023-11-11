package mural

import (
	"fmt"
	"html/template"
	"mural/controller/mural/service"
	"mural/controller/shared"
	"mural/model"
	"os"
	"strings"
	"time"

	"github.com/ryanbradynd05/go-tmdb"
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

func getReleaseYear(movie tmdb.MovieShort) string {
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

func getCurrentTheme() string {
	return service.GetCurrentDecade()
}

func getVersion() string {
	return os.Getenv("VERSION")
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

type InfoButton struct {
	Button shared.Button
	Session model.Session
}

func newInfoButton(
	text string,
	game model.Session,
) ShareButton {
	return ShareButton{
		Button: shared.Button{
			Text: text,
			Disabled: false,
		},
		Session: game,
	}
}


type ShareButton struct {
	Button shared.Button
	Session model.Session
}

func newShareButton(
	text string,
	disabled bool,
	game model.Session,
) ShareButton {
	return ShareButton{
		Button: shared.Button{
			Text: text,
			Disabled: disabled,
		},
		Session: game,
	}
}

type StatsButton struct {
	Button shared.Button
	Session model.Session
}

func newStatsButton(
	text string,
	game model.Session,
) StatsButton {
	return StatsButton{
		Button: shared.Button{
			Text: text,
			Disabled: false,
		},
		Session: game,
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