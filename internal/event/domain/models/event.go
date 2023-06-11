package models

import (
	"time"

	util "github.com/dafailyasa/event-micro/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
	ID          primitive.ObjectID `json:"id" bson:"_id", omitempty`
	Title       string             `json:"title" bson:"title"`
	Images      []string           `json:"images" bson:"images"`
	Description string             `json:"description" bson:"description"`
	IsActive    *bool              `json:"isActive" bson:"isActive"`
	Status      util.StatusTicket  `json"status" bson:"status"`
	StartDate   time.Time          `json:"startDate" bson:"startDate"`
	EndDate     time.Time          `json:"endDate" bson:"endDate"`
	Creator     Creator            `json:"creator" bson:"creator", omitempty`
}

type Creator struct {
	Avatar string `bson:"avatar"`
	Name   string `bson:"name"`
}
