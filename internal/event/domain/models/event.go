package models

import (
	"time"

	customErr "github.com/dafailyasa/event-micro/pkg/custom-error"
	util "github.com/dafailyasa/event-micro/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Images      []string           `json:"images" bson:"images"`
	Description string             `json:"description" bson:"description"`
	IsActive    *bool              `json:"isActive" bson:"isActive"`
	Status      util.StatusTicket  `json:"status" bson:"status"`
	StartDate   time.Time          `json:"startDate" bson:"startDate"`
	EndDate     time.Time          `json:"endDate" bson:"endDate"`
	Creator     Creator            `json:"creator" bson:"creator"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time          `json:"UpdatedAt" bson:"UpdatedAt"`
}

type Creator struct {
	Avatar string `bson:"avatar"`
	Name   string `bson:"name"`
}

func (e *Event) Validate() error {
	var errors error
	if !util.IsValidStringWithLength(e.Title, 5) {
		errors = customErr.ErrTitleEvent
	}

	if !util.IsValidStringWithLength(e.Description, 10) {
		errors = customErr.ErrDescEvent
	}

	if errors != nil {
		return errors
	}
	return nil
}
