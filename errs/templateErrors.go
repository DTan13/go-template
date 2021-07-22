package errs

import (
	"net/http"
)

var TemplateErrorNil *TemplateError = nil

type TemplateError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e TemplateError) Error() string {
	return e.Message
}

func NewTemplateServerError(message string) *TemplateError {
	return &TemplateError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
