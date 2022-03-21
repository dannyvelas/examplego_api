package api

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Error().Msgf("Error encoding response: %s", err)

		if _, err := io.WriteString(w, ErrInternalServerError.Error()); err != nil {
			log.Error().Msgf("Error sending Internal Server Error response: %q", err)
		}
	}
}

func RespondError(w http.ResponseWriter, internalErr error, apiErr APIError) {
	statusCode, message := apiErr.APIError()
	if statusCode == http.StatusInternalServerError {
		log.Error().Msg(internalErr.Error())
	} else {
		log.Debug().Msg(internalErr.Error())
	}
	RespondJSON(w, statusCode, message)
}
