package api

import (
	"fmt"
	"github.com/dannyvelas/examplego_api/storage"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	"net/http"
)

func ReviewsRouter(reviewsRepo storage.ReviewsRepo) func(chi.Router) {
	return func(r chi.Router) {
		r.Get("/active", getActive(reviewsRepo))
		r.Get("/all", getAll(reviewsRepo))
	}
}

func getActive(reviewsRepo storage.ReviewsRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("Get Active Reviews Endpoint")

		size := toUint(r.URL.Query().Get("size"))
		page := toUint(r.URL.Query().Get("page"))
		boundedSize, offset := getBoundedSizeAndOffset(size, page)

		activeReviews, err := reviewsRepo.GetActive(boundedSize, offset)
		if err != nil {
			err := fmt.Errorf("reviews_router: GetActive: Error querying reviewsRepo: %v", err)
			respondError(w, err, errInternalServerError)
			return
		}

		respondJSON(w, http.StatusOK, activeReviews)
	}
}

func getAll(reviewsRepo storage.ReviewsRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("Get All Endpoint")

		size := toUint(r.URL.Query().Get("size"))
		page := toUint(r.URL.Query().Get("page"))
		boundedSize, offset := getBoundedSizeAndOffset(size, page)

		allReviews, err := reviewsRepo.GetAll(boundedSize, offset)
		if err != nil {
			err = fmt.Errorf("reviews_router: getAll: Error querying reviewsRepo: %v", err)
			respondError(w, err, errInternalServerError)
			return
		}

		respondJSON(w, http.StatusOK, allReviews)
	}
}
