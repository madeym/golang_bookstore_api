package helper

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ValidationErrorResponse struct {
	Status  int
	Message []string
}

type SingleErrorResponse struct {
	Status  int
	Message string
}

func ShouldBindJSONError(c *gin.Context, err error) bool {
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error pada data %s, error: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, ValidationErrorResponse{
			Status:  http.StatusBadRequest,
			Message: errorMessages,
		})
		return true
	}
	return false
}

func ResponseError(c *gin.Context, msg string, err error) {
	var errorMessage string
	var httpstatus int
	if msg != "" {
		errorMessage = msg
		httpstatus = http.StatusBadRequest
	}
	if msg == "" {
		errorMessage = err.Error()
		httpstatus = http.StatusInternalServerError
	}
	c.JSON(httpstatus, SingleErrorResponse{
		Status:  httpstatus,
		Message: errorMessage,
	})
	return
}
