package controller

import "html/template"

type TemplateController struct {
	Template *template.Template
	Name string
}