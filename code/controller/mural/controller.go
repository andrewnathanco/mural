package mural

import (
	"html/template"
	controller_template "mural/controller/mural/template"
	"mural/controller/shared"
	"mural/model"
)


type MuralController struct { }

func NewMuralController() (*MuralController) {
	return &MuralController{}
}

func (mc MuralController) GetTemplates() []model.TemplateController {
	func_map := template.FuncMap{
        "mod": mod,
		"newButton": shared.NewButton,
		"newFlipButton": newFlipButton,
		"newShareButton": newShareButton,
		"newSelectItem": newSelectItem,
		"newSelectTile": newSelectTile,
		"newStatsButton": newStatsButton,
		"getVersion": getVersion,
    }
	

	templates := []model.TemplateController{
		controller_template.NewMuralTemplateController(func_map),
		controller_template.NewErrorTemplateController(),
		controller_template.NewSelectedTilesController(func_map), 
		controller_template.NewFlippedTileController(func_map),
		controller_template.NewAnswerTemplate(func_map), 
		controller_template.NewDialogTempalte(func_map),
	}

	return templates
}
