package apperror

import (
	"encoding/json"
	"fmt"
)

var (
	ErrNotFound = NewAppError("not found", "NS-000010", "")
)

type AppError struct {
	Err              error  `json:"-"`
	Message          string `json:"message,omitempty"`
	DeveloperMessage string `json:"developer_message,omitempty"`
	Code             string `json:"code,omitempty"`
}

func NewAppError(message, code, developerMessage string) *AppError {
	return &AppError{
		Err:              fmt.Errorf(message),
		Code:             code,
		Message:          message,
		DeveloperMessage: developerMessage,
	}
}

func (e *AppError) Error() string {
	return e.Err.Error()
}

func (e *AppError) Unwrap() error { return e.Err }

func (e *AppError) Marshal() []byte {
	bytes, err := json.Marshal(e)
	if err != nil {
		return []byte(fmt.Sprintf(`{"message": "internal error", "developer_message": "%s"}`, err.Error()))
	}
	return bytes
}

func UnauthorizedError(message string) *AppError {
	return NewAppError(message, "NS-000003", "unauthorized access")
}

func BadRequestError(message string) *AppError {
	return NewAppError(message, "NS-000002", "invalid request data")
}

func systemError(developerMessage string) *AppError {
	return NewAppError("system error", "NS-000001", developerMessage)
}

func APIError(code, message, developerMessage string) *AppError {
	return NewAppError(message, code, developerMessage)
}
