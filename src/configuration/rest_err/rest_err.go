package rest_err

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad requests",
		Code:    http.StatusBadRequest,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Not Found",
		Code:    http.StatusNotFound,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad requests",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Internal server error",
		Code:    http.StatusInternalServerError,
	}
}
