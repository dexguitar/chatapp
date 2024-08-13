package utils

import (
	"errors"
	"strings"
)

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidCreds = errors.New("invalid username or password")

type CustomError struct {
	Message    string
	StatusCode int
	Err        error
}

func (e CustomError) Error() string {
	return e.Err.Error()
}

func NewCustomError(message string, statusCode int, err error) CustomError {
	return CustomError{
		Message:    message,
		StatusCode: statusCode,
		Err:        err,
	}
}

func IsNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "user not found")
}
