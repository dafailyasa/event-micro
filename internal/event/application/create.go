package application

import (
	"time"

	"github.com/dafailyasa/event-micro/internal/event/domain/models"
)

func (app *EventApp) Create(createReq *models.CreateEventRequest) error {
	event := &models.Event{
		Title:       createReq.Title,
		Images:      createReq.Images,
		Description: createReq.Description,
		IsActive:    createReq.IsActive,
		Status:      createReq.Status,
		StartDate:   createReq.StartDate,
		EndDate:     createReq.EndDate,
		CreatedAt:   time.Now(),
	}

	if err := app.validator.Struct(event); err != nil {
		return err
	}

	err := app.repo.Save(event)
	if err != nil {
		app.logger.Error("Failed to save event", err)
		return err
	}

	app.logger.Info("Success create event", event)
	return nil
}
