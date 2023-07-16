package handlers

import (
	"fmt"

	util "github.com/dafailyasa/event-micro/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func (hdl *EventHdl) FindById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	event, err := hdl.app.FindById(id)

	if err != nil {
		hdl.logger.Error(err.Error())
		if err == mongo.ErrNoDocuments {
			msg := fmt.Sprintf("Event with %s not found", id)
			return ctx.Status(fiber.StatusNotFound).JSON(
				util.FailedRes(msg, fiber.StatusNotFound, nil),
			)
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(
			util.FailedRes(err.Error(), fiber.StatusBadRequest, nil),
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		util.SuccessRes(event, fiber.StatusOK),
	)
}
