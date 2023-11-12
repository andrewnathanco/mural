package worker

import (
	"log/slog"
)

type TMDBWorker struct { }

func NewTMDBWorker() TMDBWorker{
	return TMDBWorker{}
}

func (tw TMDBWorker) CacheAnswers() {
	slog.Info("Caching Answers From TMDB")
	// movie_controller := movie.NewTMDBController()

	// // get current page
	// current_page, err := db.DAL.GetCurrentMoviePageFromDB()
	// if err != nil {
	// 	slog.Error(fmt.Errorf("could not get current movie page: %w", err).Error())
	// 	return
	// }

	// decades := []string{
	// 	"2020",
	// 	"2010",
	// 	"2000",
	// 	"1990",
	// 	"1980",
	// 	"1970",
	// 	"1960",
	// 	"1950",
	// 	"1940",
	// }

	// for _, decade := range decades {
	// 	answers, err := movie_controller.GetAnswersByDecade(current_page + 1, decade)
	// 	if err != nil {
	// 		slog.Error(fmt.Errorf("could not get answers: %w", err).Error())
	// 		return
	// 	}

	// 	// cache answers
	// 	err = db.DAL.CacheAnswersInDatabase(answers)
	// 	if err != nil {
	// 		slog.Error(fmt.Errorf("could not cache answers: %w", err).Error())
	// 		return
	// 	}
	// }

	// // set next page
	// err = db.DAL.SetCurrentMoviePageFromDB()
	// if err != nil {
	// 	slog.Error(fmt.Errorf("could not set current movie page: %w", err).Error())
	// 	return
	// }
}