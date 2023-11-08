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
		"view/mural/game/input/answer-input.html",
		"view/mural/game/input/answer-input.tmpl",
		"view/mural/game/input/answer-options.html",
		"view/mural/game/input/answer-options.tmpl",
	}
)

func NewAnswersTemplateController(func_map template.FuncMap) model.TemplateController {
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


func NewAnswerInputTemplateController(func_map template.FuncMap) model.TemplateController {
	answer_template_files := []string{}

	// answers
	answer_template_files = append(answer_template_files, AnswerTemplates...)

	// buttons
	answer_template_files = append(answer_template_files, ButtonTemplates...)

	answer := template.Must(
		template.New("answer-input").Funcs(func_map).
		ParseFiles(
			answer_template_files...
		),
	)

	return model.TemplateController{
		Template: answer,
		Name: "answer-input.html",
	}
}

func NewAnswerOptionsTemplateController(func_map template.FuncMap) model.TemplateController {
	answer_template_files := []string{}

	// answers
	answer_template_files = append(answer_template_files, AnswerTemplates...)

	// buttons
	answer_template_files = append(answer_template_files, ButtonTemplates...)

	answer := template.Must(
		template.New("answer-options").Funcs(func_map).
		ParseFiles(
			answer_template_files...
		),
	)

	return model.TemplateController{
		Template: answer,
		Name: "answer-options.html",
	}
}
