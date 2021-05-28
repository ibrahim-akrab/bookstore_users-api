package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   error  `json:"error"`
}

func NewInternalServerError(err error) *RestErr {
	return &RestErr{
		Message: http.StatusText(http.StatusInternalServerError),
		Status:  http.StatusInternalServerError,
		Error:   err,
	}
}

func NewBadRequestError(err error) *RestErr {
	return &RestErr{
		Message: http.StatusText(http.StatusBadRequest),
		Status:  http.StatusBadRequest,
		Error:   err,
	}
}

func NewNotFoundError(err error) *RestErr {
	return &RestErr{
		Message: http.StatusText(http.StatusNotFound),
		Status:  http.StatusNotFound,
		Error:   err,
	}
}
