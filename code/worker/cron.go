package worker

import (
	"log/slog"
	"mural/app"
	"mural/controller/movie"
	"time"

	"github.com/go-co-op/gocron"
)

type MuralScheduler struct {
	Scheduler    *gocron.Scheduler
	MuralService app.MuralService
	MuralWorker  MuralWorker
	TMDBWorker   TMDBWorker
}

func NewMuralSchedular(
	tmdb_controller movie.TMDBController,
	mural_service app.MuralService,
) *MuralScheduler {
	mural_worker := MuralWorker{
		MuralService: mural_service,
	}
	tmdb_worker := TMDBWorker{
		controller: tmdb_controller,
	}
	return &MuralScheduler{
		Scheduler:    gocron.NewScheduler(time.UTC),
		MuralWorker:  mural_worker,
		TMDBWorker:   tmdb_worker,
		MuralService: mural_service,
	}

}

func (ms MuralScheduler) StartScheduler() {
	slog.Info("Starting the schedular")
	ms.Scheduler.StartAsync()
}
