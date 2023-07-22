package ports

import "github.com/gofiber/fiber/v2"

type EventHandlers interface {
	Create(ctx *fiber.Ctx) error
	FindById(ctx *fiber.Ctx) error
	FindEventWithPagination(cxt *fiber.Ctx) error
}
