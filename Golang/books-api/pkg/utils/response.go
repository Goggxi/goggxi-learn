package utils

import (
	"github.com/gofiber/fiber/v2"
)

type JSONResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func RespondWithError(c *fiber.Ctx, statusCode int, message string, errors interface{}) error {
	response := JSONResponse{
		Message: message,
		Errors:  errors,
	}
	return c.Status(statusCode).JSON(response)
}

func RespondWithSuccess(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := JSONResponse{
		Message: message,
		Data:    data,
	}
	return c.Status(statusCode).JSON(response)
}
