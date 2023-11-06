package template

import (
	"html/template"
	"mural/model"
)

var (
	MuralTemplate = []string{
		"view/mural/mural.html", 
		"view/mural/mural.tmpl", 
	}

	GameBoardTemplate = []string {
		"view/mural/game/board/game-board.html",
		"view/mural/game/board/game-board.tmpl",
		"view/mural/game/board/penalty-board.tmpl",
	}

)

func NewMuralTemplateController(
	func_map template.FuncMap,
) model.TemplateController {
	mural_template_files := []string{}


	// add buttons
	mural_template_files = append(mural_template_files, MuralTemplate...)

	// add buttons
	mural_template_files = append(mural_template_files, MuralTemplate...)

	// add buttons
	mural_template_files = append(mural_template_files, ButtonTemplates...)

	// add game board
	mural_template_files = append(mural_template_files, GameBoardTemplate...)

	// add answers
	mural_template_files = append(mural_template_files, AnswerTemplates...)

	// add tiles
	mural_template_files = append(mural_template_files, TilesTemplate...)

	mural_template := template.Must(
		template.New("mural_template").Funcs(func_map).ParseFiles(
			mural_template_files...
		))

	return model.TemplateController{
		Template: mural_template,
		Name: "mural.html",
	}

}
