package errors

type ApiError struct {
	ErrorCode string      `json:"error_code"`
	Status    int         `json:"status"`
	Message   string      `json:"message"`
	Metadata  interface{} `json:"metadata"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func NewApiError(err BaseError) ApiError {
	return ApiError{
		ErrorCode: err.ErrorCode,
		Message:   err.Message,
		Metadata:  err.Metadata,
		Status:    err.Status,
	}
}
