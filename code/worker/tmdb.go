package worker

import (
	"fmt"
	"log/slog"
	"mural/app"
	"mural/config"
	"mural/controller/movie"
)

type TMDBWorker struct {
	controller movie.TMDBController
}

func NewTMDBWorker() TMDBWorker {
	return TMDBWorker{}
}

func (tw TMDBWorker) CacheAnswers(
	service app.MuralService,
) {
	slog.Info("Caching Answers From TMDB")
	for _, decade := range config.DecadeOptions {
		movies, err := tw.controller.GetMoviesByDecade(service.Meta.LastTMDBMoviePage+1, decade)
		if err != nil {
			slog.Error(fmt.Errorf("could not get answers: %w", err).Error())
			return
		}

		// cache answers
		err = service.DAL.SaveMovies(movies)
		if err != nil {
			slog.Error(fmt.Errorf("could not cache answers: %w", err).Error())
			return
		}
	}

	service.Meta.LastTMDBMoviePage = service.Meta.LastTMDBMoviePage + 1
	config.Must(service.DAL.UpsertMeta(service.Meta))
}
