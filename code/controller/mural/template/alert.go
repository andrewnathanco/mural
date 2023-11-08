package template

import (
	"html/template"
	"mural/model"
)

var (
	AlertTemplates = []string{
		"view/mural/alerts/copied-failure.html",
		"view/mural/alerts/copied-success.html",
		"view/mural/alerts/need-answer.html",
	}
)


func NewCopiedAlertFailureTemplateController(func_map template.FuncMap) model.TemplateController {
	alert_template_files := []string{}

	// dialog
	alert_template_files = append(alert_template_files, AlertTemplates...)

	alert_template := template.Must(
		template.New("copied-alert-failure").Funcs(func_map).
		ParseFiles(
			alert_template_files...
		),
	)

	return model.TemplateController{
		Template: alert_template,
		Name: "copied-failure.html",
	}
}


func NewCopiedAlertSuccessTemplateController(func_map template.FuncMap) model.TemplateController {
	alert_template_files := []string{}

	// dialog
	alert_template_files = append(alert_template_files, AlertTemplates...)

	alert_template := template.Must(
		template.New("copied-alert-success").Funcs(func_map).
		ParseFiles(
			alert_template_files...
		),
	)

	return model.TemplateController{
		Template: alert_template,
		Name: "copied-success.html",
	}
}

func NewNeedAnswerAlertSuccessTemplateController(func_map template.FuncMap) model.TemplateController {
	alert_template_files := []string{}

	// dialog
	alert_template_files = append(alert_template_files, AlertTemplates...)

	alert_template := template.Must(
		template.New("need-answer-alert").Funcs(func_map).
		ParseFiles(
			alert_template_files...
		),
	)

	return model.TemplateController{
		Template: alert_template,
		Name: "need-answer.html",
	}
}