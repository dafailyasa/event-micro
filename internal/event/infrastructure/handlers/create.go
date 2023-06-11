package handlers

import (
	"github.com/dafailyasa/event-micro/internal/event/domain/models"
	util "github.com/dafailyasa/event-micro/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

func (hdl *EventHdl) Create(ctx *fiber.Ctx) error {
	eventReq := new(models.CreateEventRequest)
	if err := ctx.BodyParser(eventReq); err != nil {
		hdl.logger.Error(err.Error())

		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(
			util.FailedRes(err.Error(), fiber.StatusUnprocessableEntity, nil),
		)
	}

	if err := hdl.app.Create(eventReq); err != nil {
		hdl.logger.Error(err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(
			util.FailedRes(err.Error(), fiber.StatusBadRequest, nil),
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		util.SuccessRes(nil, fiber.StatusCreated),
	)
}
