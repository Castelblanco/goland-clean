package errors

type BaseError struct {
	Message   string      `json:"message"`
	ErrorCode string      `json:"error_code"`
	Status    int         `json:"status"`
	Metadata  interface{} `json:"metadata"`
}
