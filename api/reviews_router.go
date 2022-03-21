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
		r.Get("/active", GetActive(reviewsRepo))
		r.Get("/all", GetAll(reviewsRepo))
	}
}

func GetActive(reviewsRepo storage.ReviewsRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("Get Active Reviews Endpoint")

		size := ToUint(r.URL.Query().Get("size"))
		page := ToUint(r.URL.Query().Get("page"))
		boundedSize, offset := GetBoundedSizeAndOffset(size, page)

		activeReviews, err := reviewsRepo.GetActive(boundedSize, offset)
		if err != nil {
			err := fmt.Errorf("reviews_router: GetActive: Error querying reviewsRepo: %v", err)
			RespondError(w, err, ErrInternalServerError)
			return
		}

		RespondJSON(w, http.StatusOK, activeReviews)
	}
}

func GetAll(reviewsRepo storage.ReviewsRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debug().Msg("Get All Endpoint")

		size := ToUint(r.URL.Query().Get("size"))
		page := ToUint(r.URL.Query().Get("page"))
		boundedSize, offset := GetBoundedSizeAndOffset(size, page)

		allReviews, err := reviewsRepo.GetAll(boundedSize, offset)
		if err != nil {
			err = fmt.Errorf("reviews_router: GetAll: Error querying reviewsRepo: %v", err)
			RespondError(w, err, ErrInternalServerError)
			return
		}

		RespondJSON(w, http.StatusOK, allReviews)
	}
}
