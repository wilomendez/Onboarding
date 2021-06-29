package utils

import (
	"errors"
	"fmt"
)

type ErrorMessage struct {
	Type string
	Err  error
}

func (err *ErrorMessage) Error() string {
	return fmt.Sprintf("type: %s,  error: %v", err.Type, err.Err)
}

func ValidationError(msg string) error {
	return &ErrorMessage{
		Type: ValidationErrorMessage,
		Err:  errors.New(msg),
	}
}
