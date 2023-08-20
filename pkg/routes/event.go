package routes

import (
	eventHdl "github.com/dafailyasa/event-micro/internal/event/domain/ports"
	"github.com/gofiber/fiber/v2"
)

func Event(app fiber.Router, handler eventHdl.EventHandlers) {
	prefix := app.Group("event")
	prefix.Post("", handler.Create)
	prefix.Get("/search", handler.FindEventWithPagination)
	prefix.Get("/:id", handler.FindById)
	prefix.Patch("/:id", handler.Update)
}
