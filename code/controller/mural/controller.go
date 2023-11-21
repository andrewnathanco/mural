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
		"convertStringToHTML":      convertStringToHTML,
		"getReleaseYear":           getReleaseYear,
		"addCommasToNumber":        addCommaToNumber,
		"getNumberOfFlippedTiles":  getNumberOfFlippedTiples,
		"getDecadeString":          getDecadeString,
		"getHaveString":            getHaveString,
		"getSelectedTileFromBoard": getSelectedTileFromBoard,
		"getShareable":             getShareable,
		"getSimpleShare":           getSimpleShare,
	}

	templates := []model.TemplateController{
		// index
		controller_template.NewMuralTemplateController(func_map),
		controller_template.NewErrorTemplateController(func_map),

		// tiles
		controller_template.NewSelectedTilesController(func_map),
		controller_template.NewFlippedTileController(func_map),

		// dialogs
		controller_template.NewStatsDialogTemplateController(func_map),
		controller_template.NewInfoDialogTemplateController(func_map),

		// answers
		controller_template.NewAnswersTemplateController(func_map),
		controller_template.NewAnswerInputTemplateController(func_map),
		controller_template.NewAnswerOptionsTemplateController(func_map),

		// share
		controller_template.NewShareErrorTemplateController(func_map),
		controller_template.NewShareLinkTemplateController(func_map),
		controller_template.NewShareTemplateController(func_map),
		controller_template.NewShareDialogTemplateController(func_map),
	}

	return templates
}
