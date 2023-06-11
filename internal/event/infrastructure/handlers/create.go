package handlers

import (
	"github.com/dafailyasa/event-micro/internal/event/domain/models"
	"github.com/gofiber/fiber/v2"
)

func (hdl *EventHdl) Create(ctx *fiber.Ctx) error {
	eventReq := new(models.CreateEventRequest)
	if err := ctx.BodyParser(eventReq); err != nil {
		hdl.logger.Error(err.Error())
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return nil
}
