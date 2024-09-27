package errs

import (
	"errors"
)

var ErrUserNotFound = errors.New("user not found")
var ErrInvalidCreds = errors.New("invalid username or password")
var ErrInternal = errors.New("internal server error")
var ErrUserExists = errors.New("user already exists")

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
