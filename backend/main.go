package main

import (
	"log"

	"backend/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET"},
	}))

	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
