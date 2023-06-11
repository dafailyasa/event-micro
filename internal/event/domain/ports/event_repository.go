package ports

import "github.com/dafailyasa/event-micro/internal/event/domain/models"

type EventRepository interface {
	Save(event *models.Event) error
}
