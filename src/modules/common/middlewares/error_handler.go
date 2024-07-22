package middlewares

import (
	"github.com/Castelblanco/goland-clean/src/modules/common/responses/errors"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err interface{}) error {
	var apiError *errors.ApiError

	if er, ok := err.(errors.ErrorBadRequest); ok {
		apiError = &errors.ApiError{
			ErrorCode: er.ErrorCode,
			Status:    er.Status,
			Message:   er.Message,
		}
		return ctx.Status(apiError.Status).JSON(apiError)
	}

	if er, ok := err.(errors.ErrorResourceNotFound); ok {
		apiError = &errors.ApiError{
			ErrorCode: er.ErrorCode,
			Status:    er.Status,
			Message:   er.Message,
		}
		return ctx.Status(apiError.Status).JSON(apiError)
	}

	if er, ok := err.(errors.StorageError); ok {
		apiError = &errors.ApiError{
			ErrorCode: er.ErrorCode,
			Status:    er.Status,
			Message:   er.Message,
			Metadata:  er.Metadata,
		}
		return ctx.Status(apiError.Status).JSON(apiError)
	}

	return ctx.Status(fiber.StatusInternalServerError).JSON(errors.NewApiError(errors.BaseError{
		Message:   "Internal Server Error",
		ErrorCode: "Internal Server Error",
		Status:    fiber.StatusInternalServerError,
	}))
}
