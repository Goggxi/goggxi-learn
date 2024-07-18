package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

type SuccessResponse struct {
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string      `json:"message" example:"error"`
	Error   interface{} `json:"error,omitempty"`
}

func writeSuccessResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, SuccessResponse{
		Message: message,
		Data:    data,
	})
}

func writeErrorResponse(c *gin.Context, status int, message string, errors interface{}) {
	c.JSON(status, ErrorResponse{
		Message: message,
		Error:   errors,
	})
}

func WriteSuccessResponse(c *gin.Context, message string, data interface{}) {
	switch message {
	case CREATED_MESSAGE:
		writeSuccessResponse(c, http.StatusCreated, message, data)
	default:
		writeSuccessResponse(c, http.StatusOK, message, data)
	}
}

func WriteErrorResponse(c *gin.Context, message string, errors interface{}) {
	switch message {
	case BAD_REQUEST_MESSAGE:
		writeErrorResponse(c, http.StatusBadRequest, message, errors)
	case VALIDATE_ERRORS_MESSAGE:
		writeErrorResponse(c, http.StatusBadRequest, message, errors)
	default:
		if errors == pgx.ErrNoRows.Error() {
			writeErrorResponse(c, http.StatusNotFound, NOT_FOUND_MESSAGE, errors)
		} else {
			writeErrorResponse(c, http.StatusInternalServerError, INTERNAL_SERVER_ERROR_MESSAGE, errors)
		}
	}
}
