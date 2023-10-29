package repositories

import (
	"context"
	"errors"

	"github.com/dafailyasa/event-micro/internal/event/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *EventMongoDB) Update(id primitive.ObjectID, event *models.Event) (*models.Event, error) {
	opts := options.FindOneAndUpdate()
	retAfter := options.Before
	opts.ReturnDocument = &retAfter

	result := repo.collection.FindOneAndUpdate(
		context.Background(),
		bson.M{"_id": id},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "title", Value: event.Title},
				{Key: "images", Value: event.Images},
				{Key: "description", Value: event.Description},
				{Key: "isActive", Value: event.IsActive},
				{Key: "startDate", Value: event.StartDate},
				{Key: "endDate", Value: event.EndDate},
				{Key: "creator", Value: event.Creator},
				{Key: "updatedAt", Value: event.UpdatedAt},
			}},
		},
		opts,
	)

	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return nil, result.Err()
		}
		return nil, result.Err()
	}

	updatedEvent := &models.Event{}
	if err := result.Decode(updatedEvent); err != nil {
		return nil, nil
	}

	return updatedEvent, nil
}
