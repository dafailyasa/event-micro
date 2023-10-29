package handlers

import (
	"github.com/dafailyasa/event-micro/internal/event/domain/models"
	util "github.com/dafailyasa/event-micro/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (hdl *EventHdl) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	eventReq := new(models.UpdateEventRequest)
	if err := ctx.BodyParser(eventReq); err != nil {
		hdl.logger.Error(err.Error())

		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			util.FailedRes(err.Error(), fiber.StatusUnprocessableEntity, nil),
		)
	}

	event, err := hdl.app.Update(eventReq, id)

	if err != nil {
		hdl.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(
			util.FailedRes(err.Error(), fiber.StatusBadRequest, nil),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		util.SuccessRes(event, fiber.StatusOK),
	)
}
