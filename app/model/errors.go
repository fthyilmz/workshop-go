package model

import "errors"

type ErrorResponse struct {
	Messages []string `json:"messages"`
}

func (e *ErrorResponse) Message(message string) *ErrorResponse {
	e.Messages = append(e.Messages, message)

	return e
}

var (
	ErrInternalServerError = errors.New("Internal Server Error")

	ErrNotFound = errors.New("Your requested Item is not found")

	ErrConflict = errors.New("Your Item already exist")

	ErrBadParamInput = errors.New("Given Param is not valid")
)
