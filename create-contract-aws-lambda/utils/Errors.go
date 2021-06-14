package utils

import (
	"errors"
	"fmt"
)

type loadEnvVarsError struct {
	Name string
	Type string
	Err  error
}

func (e *loadEnvVarsError) Error() string {
	return fmt.Sprintf("type %s  err %v", e.Type, e.Err)
}

var (
	errAWSAccessKeyIDNotFound = loadEnvVarsError{
		Name: "AWS_ACCESS_KEY_ID",
		Type: "Env Error",
		Err:  errors.New("AWS Access Key Id not found"),
	}
	errAWSSecretAccessKeyNotFound = loadEnvVarsError{
		Name: "AWS_SECRET_ACCESS_KEY",
		Type: "Env Error",
		Err:  errors.New("AWS Secret Access Key not found"),
	}
	errAWSSessionTokenNotFound = loadEnvVarsError{
		Name: "AWS_SESSION_TOKEN",
		Type: "Env Error",
		Err:  errors.New("AWS Session Token not found"),
	}
	unknownError string = "Unknown Error"
)

func EnvError(env string) error {
	switch env {
	case errAWSAccessKeyIDNotFound.Name:
		return errAWSAccessKeyIDNotFound.Err
	case errAWSSecretAccessKeyNotFound.Name:
		return errAWSSecretAccessKeyNotFound.Err
	case errAWSSessionTokenNotFound.Name:
		return errAWSSessionTokenNotFound.Err
	default:
		return errors.New(unknownError)
	}
}
