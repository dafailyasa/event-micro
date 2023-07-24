package application

import (
	util "github.com/dafailyasa/event-micro/pkg/utils"
)

func (app *EventApp) FindByPagination(query *util.PaginationParamsStruct) (*util.PaginationResStruct, error) {
	events, eventCount, err := app.repo.FindByPagination(query)
	if err != nil {
		return nil, err
	}

	total := app.repo.Count()
	res := util.TransformPaginationRes(events, eventCount, total, query)

	return res, nil
}
