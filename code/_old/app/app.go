package app

import (
	"mural/api"
	"mural/config"
	"mural/db"
)

var ServiceContextKey = "__service"

type MuralService struct {
	Config             config.MuralConfig
	DAL                db.IDAL
	AnalyticsContoller api.IAnalyticsController
	Meta               db.MuralMeta
}

func NewMuralService(
	dal db.IDAL,
	mural_config config.MuralConfig,
	analytics_controller api.IAnalyticsController,
) (MuralService, error) {
	mural_service := MuralService{}
	// setup service
	mural_meta, err := dal.GetMeta()
	if err != nil {
		return mural_service, err
	}

	mural_service.Meta = mural_meta
	mural_service.Config = mural_config
	mural_service.AnalyticsContoller = analytics_controller
	mural_service.DAL = dal
	return mural_service, nil
}
