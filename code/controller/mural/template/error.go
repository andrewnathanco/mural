package template

import (
	"html/template"
	"mural/model"
)

var (
	ErrorTemplate = []string{
		"view/mural/mural-error.html",
	}
)
func NewErrorTemplateController(
) model.TemplateController {
	error_template := template.Must(
		template.New("mural-error").ParseFiles(ErrorTemplate...),
	)

	return model.TemplateController{
		Template: error_template,
		Name: "mural-error.html",
	}
}