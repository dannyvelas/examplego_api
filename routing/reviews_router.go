package routing

import (
	"github.com/dannyvelas/go-backend/routing/internal"
	"github.com/dannyvelas/go-backend/storage"
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
			internal.HandleInternalError(w, "Error querying reviewRepo: "+err.Error())
			return
		}

		internal.RespondJson(w, http.StatusOK, activeReviews)
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
			internal.HandleInternalError(w, "Error querying reviewRepo: "+err.Error())
			return
		}

		internal.RespondJson(w, http.StatusOK, allReviews)
	}
}
