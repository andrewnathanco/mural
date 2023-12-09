package template

import (
	"html/template"
	"mural/model"
)

var (
	DialogTemplates = []string{
		"view/mural/dialogs/share-dialog.html",
		"view/mural/dialogs/stats-dialog.html",
		"view/mural/dialogs/stats-dialog.tmpl",
		"view/mural/dialogs/info-dialog.html",
	}

	ShareLinkTemplates = []string{
		"view/mural/dialogs/share-link.html",
		"view/mural/dialogs/share-link.tmpl",
	}
)

func NewStatsDialogTemplateController(func_map template.FuncMap) model.TemplateController {
	dialog_template_files := []string{}

	// dialog
	dialog_template_files = append(dialog_template_files, DialogTemplates...)

	// buttons
	dialog_template_files = append(dialog_template_files, ButtonTemplates...)

	// scripts
	dialog_template_files = append(dialog_template_files, ScriptTemplates...)

	stats_dialog := template.Must(
		template.New("stats").Funcs(func_map).
			ParseFiles(
				dialog_template_files...,
			),
	)

	return model.TemplateController{
		Template: stats_dialog,
		Name:     "stats-dialog.html",
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

	info_dialog := template.Must(
		template.New("info").Funcs(func_map).
			ParseFiles(
				dialog_template_files...,
			),
	)

	return model.TemplateController{
		Template: info_dialog,
		Name:     "info-dialog.html",
	}

}

func NewShareDialogTemplateController(func_map template.FuncMap) model.TemplateController {
	dialog_template_files := []string{}

	// dialog
	dialog_template_files = append(dialog_template_files, DialogTemplates...)

	// buttons
	dialog_template_files = append(dialog_template_files, ButtonTemplates...)

	// game board and tiles
	dialog_template_files = append(dialog_template_files, GameBoardTemplate...)

	dialog_template_files = append(dialog_template_files, TilesTemplate...)

	share_dialog := template.Must(
		template.New("share").Funcs(func_map).
			ParseFiles(
				dialog_template_files...,
			),
	)

	return model.TemplateController{
		Template: share_dialog,
		Name:     "share-dialog.html",
	}

}

func NewShareLinkTemplateController(func_map template.FuncMap) model.TemplateController {
	share_link_files := []string{}

	// dialog
	share_link_files = append(share_link_files, ShareLinkTemplates...)

	share_link := template.Must(
		template.New("share-link").Funcs(func_map).
			ParseFiles(
				share_link_files...,
			),
	)

	return model.TemplateController{
		Template: share_link,
		Name:     "share-link.html",
	}

}
