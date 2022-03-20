package routing

import (
	"fmt"
	"github.com/dannyvelas/examplego_api/apierror"
	"github.com/dannyvelas/examplego_api/routing/internal"
	"github.com/dannyvelas/examplego_api/storage"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
)

func ReviewsRouter(reviewRepo storage.ReviewRepo) func(chi.Router) {
	return func(r chi.Router) {
		r.Get("/active", GetActive(reviewRepo))
		r.Get("/all", GetAll(reviewRepo))
	}
}

func GetActive(reviewRepo storage.ReviewRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("Get Active Endpoint")

		size := internal.ToUint(r.URL.Query().Get("size"))
		page := internal.ToUint(r.URL.Query().Get("page"))
		boundedSize, offset := internal.GetBoundedSizeAndOffset(size, page)

		activeReviews, err := reviewRepo.GetActive(boundedSize, offset)
		if err != nil {
			err := fmt.Errorf("Error in adminRepo.GetActive: %v", err)
			internal.RespondError(w, err, apierror.ErrInternalServerError)
			return
		}

		internal.RespondJSON(w, http.StatusOK, activeReviews)
	}
}

func GetAll(reviewRepo storage.ReviewRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("Get All Endpoint")

		size := internal.ToUint(r.URL.Query().Get("size"))
		page := internal.ToUint(r.URL.Query().Get("page"))
		boundedSize, offset := internal.GetBoundedSizeAndOffset(size, page)

		allReviews, err := reviewRepo.GetAll(boundedSize, offset)
		if err != nil {
			err = fmt.Errorf("Error in adminRepo.GetAll: %v", err)
			internal.RespondError(w, err, apierror.ErrInternalServerError)
			return
		}

		internal.RespondJSON(w, http.StatusOK, allReviews)
	}
}
