package ports

import (
	"github.com/dafailyasa/event-micro/internal/event/domain/models"
	util "github.com/dafailyasa/event-micro/pkg/utils"
)

type EventApplication interface {
	Create(eventRequest *models.CreateEventRequest) error
	FindById(id string) (*models.Event, error)
	FindByPagination(query *util.PaginationParamsStruct) (*[]models.Event, error)
}
