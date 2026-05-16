package handlers

import (
	"backend/data"

	"github.com/gofiber/fiber/v3"
)

func GetControls(c fiber.Ctx) error {
	return c.JSON(data.GetControls)
}
