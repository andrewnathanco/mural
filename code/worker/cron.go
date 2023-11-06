package worker

import (
	"time"

	"github.com/go-co-op/gocron"
)

type MuralScheduler struct {
	Scheduler *gocron.Scheduler
}

func NewMuralSchedular() (*MuralScheduler, error) {
	loc, err := time.LoadLocation("EST") 
	if err != nil {
		return nil, err
	}

	return &MuralScheduler{
		Scheduler: gocron.NewScheduler(loc),
	}, err
	
}

func (ms MuralScheduler) StartScheduler() {
	ms.Scheduler.StartAsync()
}