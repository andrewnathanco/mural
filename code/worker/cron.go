package worker

import (
	"time"

	"github.com/go-co-op/gocron"
)

type MuralScheduler struct {
	Scheduler *gocron.Scheduler
}

func NewMuralSchedular() *MuralScheduler {
	return &MuralScheduler{
		Scheduler: gocron.NewScheduler(time.UTC),
	}
	
}

func (ms MuralScheduler) StartScheduler() {
	ms.Scheduler.StartAsync()
}