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
        "sub": sub,
        "div": div,
		"newButton": shared.NewButton,
		"newFlipButton": newFlipButton,
		"newShareButton": newShareButton,
		"newInfoButton": newInfoButton,
		"newSelectItem": newSelectItem,
		"newSelectTile": newSelectTile,
		"newStatsButton": newStatsButton,
		"getVersion": getVersion,
		"convertStringToHTML": convertStringToHTML,
		"getReleaseYear": getReleaseYear,
		"addCommasToNumber": addCommaToNumber,
    }
	

	templates := []model.TemplateController{
		controller_template.NewMuralTemplateController(func_map),
		controller_template.NewErrorTemplateController(func_map),
		controller_template.NewSelectedTilesController(func_map), 
		controller_template.NewFlippedTileController(func_map),
		controller_template.NewAnswersTemplateController(func_map), 
		controller_template.NewStatsDialogTemplateController(func_map),
		controller_template.NewInfoDialogTemplateController(func_map),
		controller_template.NewCopiedAlertFailureTemplateController(func_map),
		controller_template.NewCopiedAlertSuccessTemplateController(func_map),
	}

	return templates
}
