package repositories

import (
	"context"

	"github.com/dafailyasa/event-micro/internal/event/domain/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (repo *EventMongoDB) FindByTitle(title string) (*models.Event, error) {
	event := new(models.Event)
	result := repo.collection.FindOne(context.Background(), bson.D{
		{Key: "title", Value: title},
	})

	if result.Err() != nil {
		return nil, result.Err()
	}

	err := result.Decode(event)
	if err != nil {
		return nil, err
	}

	return event, nil
}
