package worker

// need to do everything as utc
func (s MuralScheduler) RegisterWorkers(
) error {
	// new mural worker
	mural_worker := NewMuralWorker()

	// new tmdb worker
	tmdb_worker := NewTMDBWorker()

	// register session worker
	s.Scheduler.WaitForSchedule().Every(1).Day().At("3:59").Do(mural_worker.SetupNewGame)

	// register session worker
	s.Scheduler.WaitForSchedule().Every(1).Day().At("3:59").Do(mural_worker.ResetGameSessions)

	// register session worker
	s.Scheduler.Every(1).Day().At("2:00").Do(tmdb_worker.CacheAnswers)
	return nil
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