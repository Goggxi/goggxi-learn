package api

import (
	"book-api/internal/api/handlers"
	"book-api/internal/api/middlewares"
	"book-api/internal/repositories"
	"book-api/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4/pgxpool"
)

func SetupRoutes(app *fiber.App, pool *pgxpool.Pool) {
	// Repositories
	userRepo := repositories.NewUserRepository()

	// Services
	authService := services.NewAuthService(userRepo, pool)

	// Handlers
	authHandlers := handlers.NewAuthHandlers(authService)

	// Welcome route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Book API")
	})

	api := app.Group("/api/v1")

	// Public routes
	auth := api.Group("/auth")
	auth.Post("/signup", authHandlers.Signup)
	auth.Post("/login", authHandlers.Login)

	// Protected routes
	protectedAuth := auth.Use(middlewares.AuthMiddleware())
	protectedAuth.Post("/logout", authHandlers.Logout)
	protectedAuth.Get("/current_user", authHandlers.GetCurrentUser)
	protectedAuth.Post("/refresh_token", authHandlers.RefreshToken)
}
