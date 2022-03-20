package routing

import (
	"errors"
	"github.com/dannyvelas/examplego_api/apierror"
	"github.com/dannyvelas/examplego_api/routing/internal"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
)

func HelloRouter() func(chi.Router) {
	return func(r chi.Router) {
		r.Get("/", sayHello())
	}
}

func sayHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("Hello Endpoint")
		ctx := r.Context()

		userId := ctx.Value("id")
		if userId == nil {
			err := errors.New("key id not found in context")
			internal.RespondError(w, err, apierror.InternalServerError)
			return
		}

		userIdString, ok := userId.(string)
		if !ok {
			err := errors.New("key id is not string")
			internal.RespondError(w, err, apierror.InternalServerError)
			return
		}

		internal.RespondJSON(w, http.StatusOK, "hello, "+userIdString)
	}
}
