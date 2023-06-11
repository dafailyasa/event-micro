package repositories

import (
	"context"

	"github.com/dafailyasa/event-micro/internal/event/domain/models"
)

func (repo *EventMongoDB) Save(event *models.Event) error {
	_, err := repo.collection.InsertOne(context.Background(), event)
	if err != nil {
		repo.logger.Error("Error saving event", err)
		return err
	}
	return nil
}
