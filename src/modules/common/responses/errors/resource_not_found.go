package errors

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResourceNotFound struct {
	BaseError
}

func (e *ErrorResourceNotFound) Error() string {
	return e.Message
}

func NewErrorResourceNotFound(message string, metadata interface{}) ErrorResourceNotFound {
	return ErrorResourceNotFound{
		BaseError{
			Status:    fiber.StatusNotFound,
			Message:   message,
			ErrorCode: "Not Found",
			Metadata:  metadata,
		},
	}
}
