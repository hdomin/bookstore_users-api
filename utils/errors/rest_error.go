package errors

import "net/http"

type RestErr struct {
	Messge string `json:"message"`
	Status int    `json:"code"`
	Error  string `json:"error"`
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Messge: message,
		Status: http.StatusBadRequest,
		Error:  "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Messge: message,
		Status: http.StatusBadRequest,
		Error:  "not_found",
	}
}
