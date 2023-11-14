package worker

import (
	"mural/controller/movie"
	"time"

	"github.com/go-co-op/gocron"
)

type MuralScheduler struct {
	Scheduler   *gocron.Scheduler
	MuralWorker MuralWorker
	TMDBWorker  TMDBWorker
}

func NewMuralSchedular(
	tmdb_controller movie.TMDBController,
) *MuralScheduler {
	mural_worker := MuralWorker{}
	tmdb_worker := TMDBWorker{
		controller: tmdb_controller,
	}
	return &MuralScheduler{
		Scheduler:   gocron.NewScheduler(time.UTC),
		MuralWorker: mural_worker,
		TMDBWorker:  tmdb_worker,
	}

}

func (ms MuralScheduler) StartScheduler() {
	ms.Scheduler.StartAsync()
}
