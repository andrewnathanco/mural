package worker

import (
	"mural/config"
)

// need to do everything as utc
func (s MuralScheduler) RegisterWorkers() error {

	// register session worker
	s.Scheduler.WaitForSchedule().Every(1).Day().At("5:00").Do(s.MuralWorker.SetupNewGame)

	// tmdb can't go past 500 so we don't need to cache anymore
	if s.MuralService.Meta.LastTMDBMoviePage < 500 {
		s.Scheduler.Every(1).Minute().Do(s.TMDBWorker.CacheAnswers)
	}

	return nil
}

// need to do everything as utc
func (s MuralScheduler) RegisterWorkersDev() error {

	// register session worker
	s.Scheduler.StartImmediately().Every(2).Minute().Do(s.MuralWorker.SetupNewGame)

	// tmdb can't go past 500 so we don't need to cache anymore
	if s.MuralService.Meta.LastTMDBMoviePage < 500 {
		s.Scheduler.Every(1).Minute().Do(s.TMDBWorker.CacheAnswers)
	}

	return nil
}

// if any of this fails, kill the process
func (s MuralScheduler) InitProgram() {
	// tmdb can't go past 500 so we don't need to cache anymore
	if s.MuralWorker.MuralService.Meta.LastTMDBMoviePage < 500 {
		s.TMDBWorker.CacheAnswers()
	}

	// need to populate tiles
	config.Must(s.MuralWorker.MuralService.DAL.PopulateTiles(s.MuralWorker.MuralService.Config.BoardWidth))

	// select options

	// this will create our new game for us
	_, err := s.MuralService.DAL.GetCurrentGame(s.MuralService.Config)
	config.Must(err)
}
