package template

import (
	"html/template"
	"mural/model"
)

var (
	DialogTemplates = []string{
		"view/mural/dialogs/share-dialog.html",
	}
)


func NewDialogTemplateController(func_map template.FuncMap) model.TemplateController {
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
		Name: "share-dialog.html",
	}

}