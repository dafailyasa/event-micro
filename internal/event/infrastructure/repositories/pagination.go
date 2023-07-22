package repositories

import (
	"context"

	"github.com/dafailyasa/event-micro/internal/event/domain/models"
	util "github.com/dafailyasa/event-micro/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *EventMongoDB) FindByPagination(query *util.PaginationParamsStruct) (*[]models.Event, error) {
	var events []models.Event

	opts := options.Find()
	opts.SetLimit(int64(query.Limit))
	opts.SetSkip(int64(query.Skip))

	filter := bson.M{}
	if query.Search != "" {
		filter["$or"] = []bson.M{
			{
				"title": primitive.Regex{
					Pattern: query.Search,
					Options: "i",
				},
			},
			{
				"description": primitive.Regex{
					Pattern: query.Search,
					Options: "i",
				},
			},
		}
	}

	ctx := context.Background()
	cur, err := repo.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx) // Don't forget to close the cursor when done.

	for cur.Next(ctx) {
		var event models.Event
		err := cur.Decode(&event)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	//total, _ := repo.collection.CountDocuments(context.Background(), filter)

	return &events, nil
}
