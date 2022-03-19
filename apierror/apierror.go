package apierror

import (
	"net/http"
)

type APIError interface {
	APIError() (int, string)
}

type apiError struct {
	statusCode int
	message    string
}

var (
	ErrUnauthorized = apiError{http.StatusUnauthorized, "Unauthorized"}
	ErrBadRequest   = apiError{http.StatusBadRequest, "Bad Request"}
	ErrNotFound     = apiError{http.StatusNotFound, "Not Found"}
)

func (apiError apiError) Error() string {
	return apiError.message
}

func (apiError apiError) APIError() (int, string) {
	return apiError.statusCode, apiError.message
}

type wrappedAPIError struct {
	error
	apiError apiError
}

func (wrappedAPIError wrappedAPIError) Is(err error) bool {
	return wrappedAPIError.apiError == err
}

func Wrap(err error, apiError apiError) error {
	return wrappedAPIError{error: err, apiError: apiError}
}

func (wrappedAPIError wrappedAPIError) APIError() (int, string) {
	return wrappedAPIError.apiError.APIError()
}
