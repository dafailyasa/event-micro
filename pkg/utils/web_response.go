package util

import "github.com/gofiber/fiber/v2"

func SuccessRes(data interface{}, code int) fiber.Map {
	return fiber.Map{
		"status": "Success",
		"code":   code,
		"data":   data,
	}
}

func SuccessPaginationRes(data interface{}, meta interface{}, code int) fiber.Map {
	return fiber.Map{
		"status": "Sucess",
		"code":   code,
		"data":   data,
		"meta":   meta,
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
