package errorhandler

import "net/http"

type ErrorHandler struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *ErrorHandler) Error() string {
	return r.Message
}

func NewErrorHandler(message, err string, code int, causes []Causes) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalError(message string) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err:     "internal_server",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *ErrorHandler {
	return &ErrorHandler{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}
