package health

import (
	"mural/model"
)


type HealthController struct { }

func NewHealthController() (*HealthController) {
	return &HealthController{}
}

func (hc HealthController) GetTemplates() []model.TemplateController {
	return []model.TemplateController{}
}
