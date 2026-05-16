package handlers

import (
	"backend/data"

	"github.com/gofiber/fiber/v3"
)

func GetFooter(c fiber.Ctx) error {
	return c.JSON(data.GetFooter)
}
