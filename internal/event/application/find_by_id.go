package application

import (
	"errors"

	"github.com/dafailyasa/event-micro/internal/event/domain/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *EventApp) FindById(id string) (*models.Event, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("Invalid ObjectID format")
	}

	event, err := app.repo.FindById(objectId)

	if err != nil {
		return nil, err
	}

	return event, nil
}
