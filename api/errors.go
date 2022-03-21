package api

import (
	"net/http"
)

type apiError interface {
	apiError() (int, string)
}

type sentinel struct {
	statusCode int
	message    string
}

var (
	errUnauthorized        = sentinel{http.StatusUnauthorized, "Unauthorized"}
	errBadRequest          = sentinel{http.StatusBadRequest, "Bad Request"}
	errNotFound            = sentinel{http.StatusNotFound, "Not Found"}
	errInternalServerError = sentinel{http.StatusInternalServerError, "Internal Server Error"}
)

func (sentinel sentinel) Error() string {
	return sentinel.message
}

func (sentinel sentinel) apiError() (int, string) {
	return sentinel.statusCode, sentinel.message
}
