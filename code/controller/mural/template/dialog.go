package template

import (
	"html/template"
	"mural/model"
)

var (
	DialogTemplates = []string{
		"view/mural/dialogs/stats-dialog.html",
		"view/mural/dialogs/info-dialog.html",
	}
)


func NewStatsDialogTemplateController(func_map template.FuncMap) model.TemplateController {
	dialog_template_files := []string{}

	// dialog
	dialog_template_files = append(dialog_template_files, DialogTemplates...)

	// buttons
	dialog_template_files = append(dialog_template_files, ButtonTemplates...)


	share_dialog := template.Must(
		template.New("share").Funcs(func_map).
		ParseFiles(
			dialog_template_files...
		),
	)

	return model.TemplateController{
		Template: share_dialog,
		Name: "stats-dialog.html",
	}

}

func NewInfoDialogTemplateController(func_map template.FuncMap) model.TemplateController {
	dialog_template_files := []string{}

	// dialog
	dialog_template_files = append(dialog_template_files, DialogTemplates...)

	// buttons
	dialog_template_files = append(dialog_template_files, ButtonTemplates...)

	// game board and tiles
	dialog_template_files = append(dialog_template_files, GameBoardTemplate...)
	dialog_template_files = append(dialog_template_files, TilesTemplate...)

	share_dialog := template.Must(
		template.New("info").Funcs(func_map).
		ParseFiles(
			dialog_template_files...
		),
	)

	return model.TemplateController{
		Template: share_dialog,
		Name: "info-dialog.html",
	}

}