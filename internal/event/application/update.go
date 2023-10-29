package application

import (
	"errors"
	"time"

	"github.com/dafailyasa/event-micro/internal/event/domain/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *EventApp) Update(createReq *models.UpdateEventRequest, id string) (*models.Event, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("Invalid ObjectID format")
	}

	event := &models.Event{
		Title:       createReq.Title,
		Images:      createReq.Images,
		Description: createReq.Description,
		IsActive:    createReq.IsActive,
		Status:      createReq.Status,
		StartDate:   createReq.StartDate,
		EndDate:     createReq.EndDate,
		UpdatedAt:   time.Now(),
	}

	if err := app.validator.Struct(event); err != nil {
		return nil, err
	}

	event, err = app.repo.Update(objectId, event)
	if err != nil {
		app.logger.Error("Failed to update event", err)
		return nil, err
	}

	return event, nil
}
