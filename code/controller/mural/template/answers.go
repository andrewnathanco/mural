package template

import (
	"html/template"
	"mural/model"
)


var (
	AnswerTemplates = []string {
		"view/mural/game/answers.tmpl",
		"view/mural/game/answers.html",
		"view/mural/game/answer/default-answer.html",
		"view/mural/game/answer/selected-answer.html",
		"view/mural/game/answer/correct-answer.html",
		"view/mural/game/answer/wrong-answer.html",
	}
)

func NewAnswerTemplate(func_map template.FuncMap) model.TemplateController {
	answer_template_files := []string{}

	// answers
	answer_template_files = append(answer_template_files, AnswerTemplates...)

	// buttons
	answer_template_files = append(answer_template_files, ButtonTemplates...)


	answer := template.Must(
		template.New("answer").Funcs(func_map).
		ParseFiles(
			answer_template_files...
		),
	)

	return model.TemplateController{
		Template: answer,
		Name: "answers.html",
	}
}
