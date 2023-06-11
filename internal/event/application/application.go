package application

import (
	"github.com/dafailyasa/event-micro/internal/event/domain/ports"
	configurator "github.com/dafailyasa/event-micro/pkg/config/domain/ports"
	logger "github.com/dafailyasa/event-micro/pkg/logger/domain/ports"
	val "github.com/dafailyasa/event-micro/pkg/validator/domain/ports"
)

type EventApp struct {
	repo         ports.EventRepository
	validator    val.ValidatorApplication
	logger       logger.LoggerApplication
	configurator configurator.ConfigApplication
}

func NewUserApp(repo ports.EventRepository, validator val.ValidatorApplication, logger logger.LoggerApplication, configurator configurator.ConfigApplication) *EventApp {
	return &EventApp{
		repo:         repo,
		validator:    validator,
		logger:       logger,
		configurator: configurator,
	}
}
