package ports

import "github.com/gofiber/fiber/v2"

type EventHandlers interface {
	Create(ctx *fiber.Ctx) error
}
