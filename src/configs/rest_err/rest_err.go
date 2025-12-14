package resterr

import "net/http"

type RestErr struct {
	Message string 	 `json:"message"`
	Code  	int    	 `json:"code"`
	Err     string 	 `json:"error"`
	Causes 	[]Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field 	string `json:"field"`
	Message string `json:"message"`
}

func(r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message string, code int, err string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Code:    code,
		Err:   	 err,
		Causes:  causes,
	}
}

func NewNadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Err:     "bad_request",
	}
}

func NewNadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusBadRequest,
		Err:   	 "bad_request",
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusInternalServerError,
		Err:     "internal_server_error",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusNotFound,
		Err:     "not_found",
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Code:    http.StatusForbidden,
		Err:     "forbidden",
	}
}   