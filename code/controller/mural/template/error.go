package template

import (
	"html/template"
	"mural/model"
)

var (
	ErrorTemplate = []string{
		"view/mural/mural-error.html",
		"view/mural/share/share-error.html",
	}
)

func NewErrorTemplateController(
	func_map template.FuncMap,
) model.TemplateController {
	error_template := template.Must(
		template.New("mural-error").Funcs(func_map).ParseFiles(ErrorTemplate...),
	)

	return model.TemplateController{
		Template: error_template,
		Name:     "mural-error.html",
	}
}

func NewShareErrorTemplateController(
	func_map template.FuncMap,
) model.TemplateController {
	error_template := template.Must(
		template.New("share-error").Funcs(func_map).ParseFiles(ErrorTemplate...),
	)

	return model.TemplateController{
		Template: error_template,
		Name:     "share-error.html",
	}
}
