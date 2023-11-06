package worker

import (
	"time"

	"github.com/go-co-op/gocron"
)

type MuralScheduler struct {
	Scheduler *gocron.Scheduler
	MuralWorker MuralWorker
	TMDBWorker TMDBWorker
}

func NewMuralSchedular() *MuralScheduler {
	mural_worker := MuralWorker{}
	tmdb_worker := TMDBWorker{}
	return &MuralScheduler{
		Scheduler: gocron.NewScheduler(time.UTC),
		MuralWorker: mural_worker,
		TMDBWorker: tmdb_worker,
	}
	
}

func (ms MuralScheduler) StartScheduler() {
	ms.Scheduler.StartAsync()
}