package errors

import "github.com/gofiber/fiber/v2"

type ErrorBadRequest struct {
	BaseError
}

func (e *ErrorBadRequest) Error() string {
	return e.Message
}

func NewErrorBadRequest(message string, metadata interface{}) ErrorBadRequest {
	return ErrorBadRequest{
		BaseError{
			Status:    fiber.StatusBadRequest,
			Message:   message,
			ErrorCode: "Bad Request",
			Metadata:  metadata,
		},
	}
}
