package domain

import (
	"fmt"
)

type ErrorCode string

const (
	errorCodeServerFailure ErrorCode = "AO-101"
	errorCodeMissingField  ErrorCode = "AO-102"
)

var (
	ErrServerFailure = &APIError{Code: errorCodeServerFailure}
	ErrMissingField  = &APIError{Code: errorCodeMissingField}
)

type APIError struct {
	Code    ErrorCode
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("api: %s", e.Message)
}

type ErrorResponse struct {
	Error ErrorDesc `json:"error"`
}

type ErrorDesc struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}
