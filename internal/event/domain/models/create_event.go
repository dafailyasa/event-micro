package models

import (
	"time"

	util "github.com/dafailyasa/event-micro/pkg/utils"
)

type CreateEventRequest struct {
	Title       string            `json:"title" validate:"required"`
	Images      []string          `json:"images" validate:"required"`
	Description string            `json:"description" validate:"required"`
	IsActive    *bool             `json:"isActive" validate:"required,boolean"`
	Status      util.StatusTicket `json:"status" validate:"required"`
	StartDate   time.Time         `json:"startDate" validate:"required"`
	EndDate     time.Time         `json:"endDate" validate:"required"`
}
