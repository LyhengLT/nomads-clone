package main

import (
	"log"
	"os"

	"backend/routes"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	app := fiber.New()

	allowOrigins := os.Getenv("CORS_ORIGIN")
	if allowOrigins == "" {
		allowOrigins = "http://localhost:5173"
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{allowOrigins},
		AllowMethods: []string{"GET"},
	}))

	routes.SetupRoutes(app)

	app.Use("/", static.New("./dist", static.Config{
		IndexNames: []string{"index.html"},
		Browse:     false,
	}))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
