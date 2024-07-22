package errors

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type StorageError struct {
	BaseError
}

func (e StorageError) Error() string {
	return fmt.Sprintf("Message: %s, ErrorCode: %s, Status: %d, Metadata: %s,", e.Message, e.ErrorCode, e.Status, e.Metadata)
}

func NewStorageError(message string, metadata interface{}) StorageError {
	return StorageError{
		BaseError{
			Status:    fiber.StatusBadRequest,
			Message:   message,
			ErrorCode: "Storage error",
			Metadata:  metadata,
		},
	}
}
