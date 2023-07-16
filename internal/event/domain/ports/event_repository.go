package ports

import (
	"github.com/dafailyasa/event-micro/internal/event/domain/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EventRepository interface {
	Save(event *models.Event) error
	FindByTitle(title string) (*models.Event, error)
	FindById(id primitive.ObjectID) (*models.Event, error)
}
