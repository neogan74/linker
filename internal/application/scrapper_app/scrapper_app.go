package scrapper_app

import (
	"time"

	"github.com/go-co-op/gocron/v2"
)

type TrackingClient interface {
	Send(url string) (time.Time, error)
}

type ScrapperApp struct {
	Scheduler       *gocron.Scheduler
	ScrapperClient  *clients.ScapperClinet
	TrackingClinets []TrackingClient
	Repository      infra.Repository
}
