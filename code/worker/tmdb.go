package worker

import (
	"fmt"
	"log/slog"
	"mural/app"
	"mural/config"
	"mural/controller/movie"
)

type TMDBWorker struct {
	Service    app.MuralService
	controller movie.TMDBController
}

func (tw TMDBWorker) CacheAnswers() {
	slog.Info("Caching Answers From TMDB")
	for _, decade := range config.DecadeOptions {
		movies, err := tw.controller.GetMoviesByDecade(tw.Service.Meta.LastTMDBMoviePage+1, decade)
		if err != nil {
			slog.Error(fmt.Errorf("could not get answers: %w", err).Error())
			return
		}

		// cache answers
		err = tw.Service.DAL.SaveMovies(movies)
		if err != nil {
			slog.Error(fmt.Errorf("could not cache answers: %w", err).Error())
			return
		}
	}

	tw.Service.Meta.LastTMDBMoviePage = tw.Service.Meta.LastTMDBMoviePage + 1
	config.Must(tw.Service.DAL.UpsertMeta(tw.Service.Meta))
}
