package handlers

import (
	"backend/data"

	"github.com/gofiber/fiber/v3"
)

func GetMeetups(c fiber.Ctx) error {
	return c.JSON(data.GetMeetups)
}
