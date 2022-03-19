package apierror

import (
	"net/http"
)

type APIError interface {
	APIError() (int, string)
}

type sentinel struct {
	statusCode int
	message    string
}

var (
	ErrUnauthorized        = sentinel{http.StatusUnauthorized, "Unauthorized"}
	ErrBadRequest          = sentinel{http.StatusBadRequest, "Bad Request"}
	ErrNotFound            = sentinel{http.StatusNotFound, "Not Found"}
	ErrInternalServerError = sentinel{http.StatusInternalServerError, "Internal Server Error"}
)

func (sentinel sentinel) Error() string {
	return sentinel.message
}

func (sentinel sentinel) APIError() (int, string) {
	return sentinel.statusCode, sentinel.message
}
