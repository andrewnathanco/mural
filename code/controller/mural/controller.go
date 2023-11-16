package mural

import (
	"html/template"
	controller_template "mural/controller/mural/template"
	"mural/controller/shared"
	"mural/model"
)

type MuralController struct{}

func NewMuralController() *MuralController {
	return &MuralController{}
}

func (mc MuralController) GetTemplates() []model.TemplateController {
	func_map := template.FuncMap{
		"mod":                      mod,
		"sub":                      sub,
		"div":                      div,
		"bang":                     bang,
		"newButton":                shared.NewButton,
		"newFlipButton":            newFlipButton,
		"newShareButton":           newShareButton,
		"newInfoButton":            newInfoButton,
		"newSelectItem":            newSelectItem,
		"newSelectTile":            newSelectTile,
		"newStatsButton":           newStatsButton,
		"getDate":                  getDate,
		"getCurrentTheme":          getCurrentTheme,
		"convertStringToHTML":      convertStringToHTML,
		"getReleaseYear":           getReleaseYear,
		"addCommasToNumber":        addCommaToNumber,
		"getNumberOfFlippedTiles":  getNumberOfFlippedTiples,
		"getDecadeString":          getDecadeString,
		"getHaveString":            getHaveString,
		"getSelectedTileFromBoard": getSelectedTileFromBoard,
		"getShareable":             getShareable,
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
		controller_template.NewNeedAnswerAlertSuccessTemplateController(func_map),
		controller_template.NewAnswerInputTemplateController(func_map),
		controller_template.NewAnswerOptionsTemplateController(func_map),
	}

	return templates
}
