package worker

import (
	"database/sql"
	"mural/db"
)

// need to do everything as utc
func (s MuralScheduler) RegisterWorkers(
) error {

	// register session worker
	s.Scheduler.WaitForSchedule().Every(1).Day().At("4:59").Do(s.MuralWorker.SetupNewGame)

	// register session worker
	s.Scheduler.WaitForSchedule().Every(1).Day().At("4:59").Do(s.MuralWorker.ResetGameSessions)

	// register session worker
	s.Scheduler.Every(1).Day().At("3:00").Do(s.TMDBWorker.CacheAnswers)
	return nil
}

func (s MuralScheduler) InitProgram() {
	// need to manually pull a few answers to start
	s.TMDBWorker.CacheAnswers()

	_, err := db.DAL.GetCurrentGameInfo()
	if err == sql.ErrNoRows {
		// if the game doesn't exist, lets set it up
		s.MuralWorker.SetupNewGame()
	}
}

func (s MuralScheduler) RegisterWorkersFreeplay(
) error {
	// new mural worker
	mural_worker := NewMuralWorker()

	// new tmdb worker
	tmdb_worker := NewTMDBWorker()

	// register session worker
	s.Scheduler.Every(1).Minute().Do(mural_worker.SetupNewGame)

	// register session worker
	s.Scheduler.Every(1).Minute().Do(mural_worker.ResetGameSessions)
	s.Scheduler.Every(1).Minute().Do(tmdb_worker.CacheAnswers)

	// register session worker
	s.Scheduler.Every(1).Minute().Do(tmdb_worker.CacheAnswers)
	return nil
}