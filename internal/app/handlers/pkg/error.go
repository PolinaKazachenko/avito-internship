package pkg

type Error struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func NewError(message string, code string) *Error {
	return &Error{
		Message: message,
		Code:    code,
	}
}

func (e *Error) Error() string {
	return e.Message
}

var (
	ErrInvalidRequestBody  = NewError("invalid request body", "INVALID_REQUEST_BODY")
	ErrNotFound            = NewError("resource not found", "NOT_FOUND")
	ErrInternalServerError = NewError("internal server error", "INTERNAL_SERVER_ERROR")
)
