package handlers

import (
	"github.com/dafailyasa/event-micro/internal/event/domain/ports"
	logger "github.com/dafailyasa/event-micro/pkg/logger/domain/ports"
)

type EventHdl struct {
	app    ports.EventApplication
	logger logger.LoggerApplication
}

var _ ports.EventHandlers = (*EventHdl)(nil)

func NewEventHdl(app ports.EventApplication, logger logger.LoggerApplication) *EventHdl {
	return &EventHdl{app: app, logger: logger}
}
