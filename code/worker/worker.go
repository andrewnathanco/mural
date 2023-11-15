package worker

import (
	"mural/app"
	"mural/config"
)

// need to do everything as utc
func (s MuralScheduler) RegisterWorkers(
	service app.MuralService,
) error {

	// register session worker
	s.Scheduler.WaitForSchedule().Every(1).Day().At("5:01").Do(s.MuralWorker.SetupNewGame)

	// tmdb can't go past 500 so we don't need to cache anymore
	if service.Meta.LastTMDBMoviePage < 500 {
		s.Scheduler.Every(1).Minute().Do(s.TMDBWorker.CacheAnswers)
	}

	return nil
}

// need to do everything as utc
func (s MuralScheduler) RegisterWorkersDev(
	service app.MuralService,
) error {

	// register session worker
	s.Scheduler.Every(1).Minute().Do(s.MuralWorker.SetupNewGame)

	// tmdb can't go past 500 so we don't need to cache anymore
	if service.Meta.LastTMDBMoviePage < 500 {
		s.Scheduler.Every(1).Minute().Do(s.TMDBWorker.CacheAnswers)
	}

	return nil
}

// if any of this fails, kill the process
func (s MuralScheduler) InitProgram(
	service app.MuralService,
) {
	// tmdb can't go past 500 so we don't need to cache anymore
	if service.Meta.LastTMDBMoviePage < 500 {
		s.TMDBWorker.CacheAnswers(
			service,
		)
	}

	// need to populate tiles
	config.Must(service.DAL.PopulateTiles(service.Config.BoardWidth))

	// select options

	// this will create our new game for us
	_, err := service.DAL.GetCurrentGame(service.Config)
	config.Must(err)
}
