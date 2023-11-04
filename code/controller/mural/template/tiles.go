package template

import (
	"html/template"
	"mural/model"
)

var (
	TilesTemplate = []string {
		"view/mural/game/tile/default-tile.html",
		"view/mural/game/tile/penalty-tile.html",
		"view/mural/game/tile/selected/selected-tile.tmpl",
		"view/mural/game/tile/flipped-tile.html",
	}
)

func NewSelectedTilesController(func_map template.FuncMap) model.TemplateController {
	selected_template_files := []string{}

	// game board
	selected_template_files = append(selected_template_files, GameBoardTemplate...)

	// tiles
	selected_template_files = append(selected_template_files, TilesTemplate...)

	// buttons
	selected_template_files = append(selected_template_files, ButtonTemplates...)

	selected_tile := template.Must(
		template.New("selected").Funcs(func_map).
		ParseFiles(
			selected_template_files...
		),
	)

	return model.TemplateController{
		Template: selected_tile,
		Name: "game-board.html",
	}
}


func NewFlippedTileController(func_map template.FuncMap) model.TemplateController {
	flipped_template_files := []string{}

	// answers
	flipped_template_files = append(flipped_template_files, AnswerTemplates...)

	// game board
	flipped_template_files = append(flipped_template_files, GameBoardTemplate...)

	// tiles
	flipped_template_files = append(flipped_template_files, TilesTemplate...)

	// buttons
	flipped_template_files = append(flipped_template_files, ButtonTemplates...)


	flipped_tile := template.Must(
		template.New("flipped").Funcs(func_map).
		ParseFiles(
			flipped_template_files...
		),
	)

	return model.TemplateController{
		Template: flipped_tile,
		Name: "game-board.html",
	}
}