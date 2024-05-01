package goapplib

import (
	"fmt"
)

type StatusError struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func NewStatusError(code Code, message string) *StatusError {
	return &StatusError{
		Code:    code,
		Message: message,
	}
}

func NewStatusErrorf(code Code, format string, args ...interface{}) *StatusError {
	return &StatusError{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}

func NewStatusErrorFromError(code Code, err error) *StatusError {
	return &StatusError{
		Code:    code,
		Message: err.Error(),
	}
}

func ErrorAsStatusError(err error) (*StatusError, bool) {
	if err == nil {
		return nil, false
	}
	e, ok := err.(*StatusError)
	return e, ok
}
