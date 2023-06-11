package ports

import "github.com/dafailyasa/event-micro/internal/event/domain/models"

type EventApplication interface {
	Create(registerRequest *models.CreateEventRequest) error
}
