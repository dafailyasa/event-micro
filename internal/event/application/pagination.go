package application

import (
	"github.com/dafailyasa/event-micro/internal/event/domain/models"
	util "github.com/dafailyasa/event-micro/pkg/utils"
)

func (app *EventApp) FindByPagination(query *util.PaginationParamsStruct) (*[]models.Event, error) {
	events, err := app.repo.FindByPagination(query)
	if err != nil {
		return nil, err
	}

	return events, nil
}
