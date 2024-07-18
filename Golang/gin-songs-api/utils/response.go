package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"net/http"
)

type AppResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func writeResponse(c *gin.Context, status int, message string, data interface{}, errors interface{}) {
	c.JSON(status, AppResponse{
		Message: message,
		Data:    data,
		Error:   errors,
	})
}

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	switch message {
	case CREATED_MESSAGE:
		writeResponse(c, http.StatusCreated, message, data, nil)
	default:
		writeResponse(c, http.StatusOK, message, data, nil)
	}
}

func ErrorResponse(c *gin.Context, message string, errors interface{}) {
	switch message {
	case BAD_REQUEST_MESSAGE:
		writeResponse(c, http.StatusBadRequest, message, nil, errors)
	case VALIDATE_ERRORS_MESSAGE:
		writeResponse(c, http.StatusBadRequest, message, nil, errors)
	default:
		if errors == pgx.ErrNoRows.Error() {
			writeResponse(c, http.StatusNotFound, NOT_FOUND_MESSAGE, nil, errors)
		} else {
			writeResponse(c, http.StatusInternalServerError, INTERNAL_SERVER_ERROR_MESSAGE, nil, errors)
		}
	}
}
