package internal

import (
	"encoding/json"
	"errors"
	"github.com/dannyvelas/examplego_api/apierror"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

const (
	internalServerErrorResponse = "Internal Server Error"
)

func RespondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		log.Error().Msgf("Error encoding response: %s", err)

		if _, err := io.WriteString(w, internalServerErrorResponse); err != nil {
			log.Error().Msgf("Error sending Internal Server Error response: %q", err)
		}
	}
}

func RespondError(w http.ResponseWriter, err error) {
	var apiErr apierror.APIError
	if errors.As(err, &apiErr) {
		statusCode, msg := apiErr.APIError()
		RespondJSON(w, statusCode, msg)
	} else {
		log.Error().Msg(err.Error())
		RespondJSON(w, http.StatusInternalServerError, internalServerErrorResponse)
	}
}
