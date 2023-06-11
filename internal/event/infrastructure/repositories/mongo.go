package repositories

import (
	"github.com/dafailyasa/event-micro/internal/event/domain/ports"
	config "github.com/dafailyasa/event-micro/pkg/config/domain/ports"
	logger "github.com/dafailyasa/event-micro/pkg/logger/domain/ports"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventMongoDB struct {
	client       *mongo.Client
	database     *mongo.Database
	collection   *mongo.Collection
	configurator config.ConfigApplication
	logger       logger.LoggerApplication
}

var _ ports.EventRepository = (*EventMongoDB)(nil)

func NewEventMongoDB(config config.ConfigApplication, logger logger.LoggerApplication, client *mongo.Client) *EventMongoDB {
	cfg, err := config.GetConfig()
	if err != nil {
		logger.Error("Error getting configurator", err)
		return nil
	}

	return &EventMongoDB{
		client:       client,
		database:     client.Database(cfg.Database.Name),
		collection:   client.Database(cfg.Database.Name).Collection("events"),
		configurator: config,
		logger:       logger,
	}
}
