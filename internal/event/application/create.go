package application

import (
	"errors"
	"fmt"
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

	exist, err := app.repo.FindByTitle(event.Title)

	if exist != nil {
		msg := fmt.Sprintf("Title %s already used", event.Title)
		return errors.New(msg)
	}

	err = app.repo.Save(event)
	if err != nil {
		app.logger.Error("Failed to save event", err)
		return err
	}

	return nil
}
