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
	Unauthorized        = sentinel{http.StatusUnauthorized, "Unauthorized"}
	BadRequest          = sentinel{http.StatusBadRequest, "Bad Request"}
	NotFound            = sentinel{http.StatusNotFound, "Not Found"}
	InternalServerError = sentinel{http.StatusInternalServerError, "Internal Server Error"}
)

func (sentinel sentinel) Error() string {
	return sentinel.message
}

func (sentinel sentinel) APIError() (int, string) {
	return sentinel.statusCode, sentinel.message
}
