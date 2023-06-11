package util

import "github.com/gofiber/fiber/v2"

func SuccessRes(data interface{}, code int) fiber.Map {
	return fiber.Map{
		"status": "Success",
		"code":   code,
		"data":   data,
	}
}

func FailedRes(err interface{}, code int, data interface{}) fiber.Map {
	return fiber.Map{
		"status": "Failed",
		"code":   code,
		"error":  err,
		"data":   data,
	}
}
