package controllers

import (
	"fmt"
	"net/http"
)

type HandleError struct {
	StatusCode int    `json:"-"`
	Type       string `json:"type"`
	Message    string `json:"message"`
}

func (h HandleError) Error() string {
	return fmt.Sprintf(`{"code": "%v", "type": "%s", "error": "%s" }`, h.StatusCode, h.Type, h.Message)
}

var (
	errEmptyData = HandleError{
		StatusCode: http.StatusBadRequest,
		Type:       "Validation Error",
		Message:    "Input validation FAILED, ID is empty",
	}
)

func HandleCustomValidationError(message string) HandleError {
	return HandleError{
		StatusCode: http.StatusBadRequest,
		Type:       "Validation Error",
		Message:    message,
	}
}

func HandleNotFoundError(message string) HandleError {
	return HandleError{
		StatusCode: http.StatusNotFound,
		Type:       "Not Found Error",
		Message:    message,
	}
}

func HandleInternalError(message string) HandleError {
	return HandleError{
		StatusCode: http.StatusInternalServerError,
		Type:       "Internal Server Error",
		Message:    message,
	}
}
