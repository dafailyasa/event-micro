package repositories

import (
	"context"

	"github.com/dafailyasa/event-micro/internal/event/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *EventMongoDB) Update(id primitive.ObjectID, event *models.UpdateEventRequest) error {

	opts := options.Update()
	opts.SetUpsert(true)

	_, err := repo.collection.UpdateByID(context.Background(), id, bson.M{}, opts)

	if err != nil {
		return err
	}

	return nil
}
