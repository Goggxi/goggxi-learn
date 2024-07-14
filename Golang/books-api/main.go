package main

import (
	"book-api/config"
	"book-api/platform/database"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	database.ConnectPostgresDB(cfg)
	defer database.DB.Close()

	// Set up Fiber app
	app := fiber.New()

	// Define routes
	setupRoutes(app)

	// Start the server
	log.Fatal(app.Listen(":3000"))
}

func setupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
