package worker

// need to do everything as utc
func (s MuralScheduler) RegisterWorkers(
) error {

	// register session worker
	s.Scheduler.WaitForSchedule().Every(1).Day().At("4:59").Do(s.MuralWorker.SetupNewGame)

	// register session worker
	s.Scheduler.WaitForSchedule().Every(1).Day().At("4:59").Do(s.MuralWorker.ResetGameSessions)

	// register session worker
	// get current page
	// current_page, err := db.DAL.GetCurrentMoviePageFromDB()
	// if err != nil {
	// 	slog.Error(fmt.Errorf("could not get current movie page: %w", err).Error())
	// 	return err
	// }

	// // tmdb can't go past 500 so we don't need to cache anymore
	// if current_page < 500 {
	// 	s.Scheduler.Every(1).Minute().Do(s.TMDBWorker.CacheAnswers)
	// }

	return nil
}

func (s MuralScheduler) InitProgram() {
	// need to manually pull a few answers to start
	// current_page, err := db.DAL.GetCurrentMoviePageFromDB()
	// if err != nil {
	// 	slog.Error(fmt.Errorf("could not get current movie page: %w", err).Error())
	// }

	// // tmdb can't go past 500 so we don't need to cache anymore
	// if current_page < 500 {
	// 	s.TMDBWorker.CacheAnswers()
	// }

	// _, err = db.DAL.GetCurrentGameInfo()
	// if err == sql.ErrNoRows {
	// 	// if the game doesn't exist, lets set it up
	// 	s.MuralWorker.SetupNewGame()
	// }
}

func (s MuralScheduler) RegisterWorkersFreeplay(
) error {
	// new mural worker
	mural_worker := NewMuralWorker()

	// new tmdb worker
	tmdb_worker := NewTMDBWorker()

	// register session worker
	s.Scheduler.Every(2).Minute().Do(mural_worker.SetupNewGame)

	// register session worker
	s.Scheduler.Every(2).Minute().Do(mural_worker.ResetGameSessions)

	// register session worker
	s.Scheduler.Every(1).Minute().Do(tmdb_worker.CacheAnswers)
	return nil
}