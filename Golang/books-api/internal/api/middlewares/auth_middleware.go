package middlewares

import (
	"book-api/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"strings"
)

// AuthMiddleware godoc
// @Summary Authenticate requests
// @Description Middleware to authenticate requests using JWT
// @Tags middleware
// @Param Authorization header string true "Bearer <token>"
// @Failure 401 {object} utils.SwaggerErrorResponse
func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return utils.RespondWithError(c, fiber.StatusUnauthorized, "Missing authorization header", nil)
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			return utils.RespondWithError(c, fiber.StatusUnauthorized, "Invalid authorization header", nil)
		}

		token := tokenParts[1]
		claims, err := utils.VerifyToken(token)
		if err != nil {
			return utils.RespondWithError(c, fiber.StatusUnauthorized, "Invalid token", err.Error())
		}

		// Set user ID in context for use in subsequent handlers
		c.Locals("userID", claims.UserID)

		return c.Next()
	}
}
