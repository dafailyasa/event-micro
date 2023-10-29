package models

import (
	"time"

	customErr "github.com/dafailyasa/event-micro/pkg/custom-error"
	util "github.com/dafailyasa/event-micro/pkg/utils"
)

type UpdateEventRequest struct {
	Title       string            `json:"title" validate:"required"`
	Images      []string          `json:"images" validate:"required"`
	Description string            `json:"description" validate:"required"`
	IsActive    *bool             `json:"isActive" validate:"required,boolean"`
	Status      util.StatusTicket `json:"status" validate:"required"`
	StartDate   time.Time         `json:"startDate" validate:"required"`
	EndDate     time.Time         `json:"endDate" validate:"required"`
}

func (e *UpdateEventRequest) Validate() []error {
	var errors []error

	if !util.IsValidStringWithLength(e.Title, 5) {
		errors = append(errors, customErr.ErrTitleEvent)
	}

	if !util.IsValidArrayWithLength(e.Images, 1) {
		errors = append(errors, customErr.ErrImagesEvent)
	}

	if !util.IsValidStringWithLength(e.Description, 10) {
		errors = append(errors, customErr.ErrDescEvent)
	}

	if len(errors) >= 1 {
		return errors
	}

	return nil
}
