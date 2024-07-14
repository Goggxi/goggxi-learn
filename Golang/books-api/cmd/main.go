package main

import (
	"book-api/internal/api"
	"book-api/internal/config"
	"book-api/platform/database"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	database.ConnectPostgresDB(cfg)
	defer database.DB.Close()

	app := fiber.New()

	api.SetupRoutes(app, database.DB)

	log.Fatal(app.Listen(cfg.ServerAddress))
}
