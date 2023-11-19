package template

import (
	"html/template"
	"mural/model"
)

var (
	ShareTemplate = []string{
		"view/mural/share/share.html",
		"view/mural/share/share.tmpl",
	}
)

func NewShareTemplateController(
	func_map template.FuncMap,
) model.TemplateController {
	share_template_files := []string{}

	// add buttons
	share_template_files = append(share_template_files, ShareTemplate...)

	// add buttons
	share_template_files = append(share_template_files, MuralTemplate...)

	// add buttons
	share_template_files = append(share_template_files, ButtonTemplates...)

	// add game board
	share_template_files = append(share_template_files, GameBoardTemplate...)

	// add answers
	share_template_files = append(share_template_files, AnswerTemplates...)

	// add tiles
	share_template_files = append(share_template_files, TilesTemplate...)

	// add tiles
	share_template_files = append(share_template_files, ScriptTemplates...)

	share_template := template.Must(
		template.New("share_template").Funcs(func_map).ParseFiles(
			share_template_files...,
		))

	return model.TemplateController{
		Template: share_template,
		Name:     "share.html",
	}

}
