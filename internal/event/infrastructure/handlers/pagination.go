package handlers

import (
	util "github.com/dafailyasa/event-micro/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (hdl *EventHdl) FindEventWithPagination(ctx *fiber.Ctx) error {
	query := util.QueryPaginationParams(ctx)
	events, err := hdl.app.FindByPagination(query)

	if err != nil {
		hdl.logger.Error(err.Error())

		return ctx.Status(fiber.StatusBadRequest).JSON(
			util.FailedRes(err.Error(), fiber.StatusBadRequest, nil),
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		util.SuccessRes(events, fiber.StatusOK),
	)
}
