package main

import (
	_ "book-api/docs"
	"book-api/internal/api"
	"book-api/internal/config"
	"book-api/platform/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"log"
)

// @title Book API
// @version 1.0
// @description This is a sample Book API server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api/v1
func main() {
	cfg := config.LoadConfig()

	database.ConnectPostgresDB(cfg)
	defer database.DB.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // or specify domains such as "http://example.com, https://example.com"
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)

	api.SetupRoutes(app, database.DB)

	log.Fatal(app.Listen(cfg.ServerAddress))
}
