package routes

import (
	"backend/handlers"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/hero", handlers.GetHero)
	api.Get("/filters", handlers.GetFilters)
	api.Get("/controls", handlers.GetControls)
	api.Get("/dropdown", handlers.GetDropdown)
	api.Get("/sidebar", handlers.GetSidebar)
	api.Get("/cities", handlers.GetCities)
	api.Get("/meetups", handlers.GetMeetups)
	api.Get("/footer", handlers.GetFooter)
	_ = api
}
