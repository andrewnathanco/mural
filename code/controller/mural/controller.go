package mural

import (
	"html/template"
	"mural/controller/shared"
	"mural/model"
)


type MuralController struct { }

func NewMuralController() (*MuralController) {
	return &MuralController{}
}

// GetTemplates() []model.TemplateController
// GetRoutes() map[string]func(c echo.Context)error

func (mc MuralController) GetTemplates() []model.TemplateController {
	func_map := template.FuncMap{
        "mod": mod,
		"newButton": shared.NewButton,
		"newFlipButton": newFlipButton,
		"newSelectItem": newSelectItem,
		"newSelectTile": newSelectTile,
		"getVersion": getVersion,
    }

	mural_template := template.Must(
		template.New("mural_template").Funcs(func_map).ParseFiles(
			"view/mural/mural.html", 
			"view/mural/mural.tmpl", 
			"view/mural/game/board/game-board.html",
			"view/mural/game/board/game-board.tmpl",
			"view/mural/game/answers.tmpl",
			"view/mural/game/answers.html",
			"view/mural/game/answer/default-answer.html",
			"view/mural/game/answer/selected-answer.html",
			"view/mural/game/answer/correct-answer.html",
			"view/mural/game/answer/wrong-answer.html",
			"view/mural/game/tile/default-tile.html",
			"view/mural/game/tile/penalty-tile.html",
			"view/mural/game/tile/selected/selected-tile.tmpl",
			"view/mural/game/tile/flipped-tile.html",
			"view/mural/buttons/flip-button.tmpl",
			"view/mural/buttons/submit-button.tmpl",
			"view/mural/buttons/share-button.tmpl",
		))



	mural_template_controller := model.TemplateController{
		Template: mural_template,
		Name: "mural.html",
	}

	error_template := template.Must(
		template.New("mural-error").ParseFiles("view/mural/mural-error.html"),
	)
	error_template_controller := model.TemplateController{
		Template: error_template,
		Name: "mural-error.html",
	}

	selected_tile := template.Must(
		template.New("selected").Funcs(func_map).
		ParseFiles(
			"view/mural/game/board/game-board.html",
			"view/mural/game/board/game-board.tmpl",
			"view/mural/game/tile/default-tile.html",
			"view/mural/game/tile/selected/selected-tile.tmpl",
			"view/mural/game/tile/flipped-tile.html",
			"view/mural/buttons/flip-button.tmpl",
			"view/mural/game/tile/penalty-tile.html",
		),
	)

	selected_tile_controller := model.TemplateController{
		Template: selected_tile,
		Name: "game-board.html",
	}

	flipped_tile := template.Must(
		template.New("flipped").Funcs(func_map).
		ParseFiles(
			"view/mural/game/answers.html",
			"view/mural/game/answers.tmpl",
			"view/mural/game/answer/default-answer.html",
			"view/mural/game/answer/selected-answer.html",
			"view/mural/game/answer/correct-answer.html",
			"view/mural/game/answer/wrong-answer.html",
			"view/mural/game/answers.html",
			"view/mural/game/answers.tmpl",
			"view/mural/buttons/submit-button.tmpl",
			"view/mural/game/board/game-board.html",
			"view/mural/game/board/game-board.tmpl",
			"view/mural/game/tile/default-tile.html",
			"view/mural/game/tile/penalty-tile.html",
			"view/mural/game/tile/selected/selected-tile.tmpl",
			"view/mural/game/tile/flipped-tile.html",
			"view/mural/buttons/flip-button.tmpl",
			"view/mural/buttons/share-button.tmpl",
		),
	)

	flipped_tile_controller := model.TemplateController{
		Template: flipped_tile,
		Name: "game-board.html",
	}

	answer := template.Must(
		template.New("answer").Funcs(func_map).
		ParseFiles(
			"view/mural/game/answers.html",
			"view/mural/game/answers.tmpl",
			"view/mural/game/answer/default-answer.html",
			"view/mural/game/answer/selected-answer.html",
			"view/mural/game/answer/correct-answer.html",
			"view/mural/game/answer/wrong-answer.html",
			"view/mural/buttons/submit-button.tmpl",
			"view/mural/buttons/share-button.tmpl",
		),
	)

	answer_controller := model.TemplateController{
		Template: answer,
		Name: "answers.html",
	}

	templates := []model.TemplateController{
		mural_template_controller, error_template_controller,
		selected_tile_controller, flipped_tile_controller,
		answer_controller,
	}
	return templates
}
