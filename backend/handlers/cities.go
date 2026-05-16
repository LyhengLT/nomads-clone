package handlers

import (
	"backend/data"

	"github.com/gofiber/fiber/v3"
)

func GetCities(c fiber.Ctx) error {
	return c.JSON(data.GetCities)
}
