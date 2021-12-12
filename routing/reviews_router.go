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
		log.Info().Msg("Get All Endpoint")

		page := internal.ToUint(r.URL.Query().Get("page"))
		size := internal.ToUint(r.URL.Query().Get("size"))
		limit, offset := internal.PagingToLimitOffset(page, size)

		activeReviews, err := reviewRepo.GetActive(limit, offset)
		if err != nil {
			internal.HandleInternalError(w, "Error querying reviewRepo: "+err.Error())
			return
		}

		internal.RespondJson(w, http.StatusOK, activeReviews)
	}
}

func GetAll(reviewRepo storage.ReviewRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg("Get All Endpoint")

		page := internal.ToUint(r.URL.Query().Get("page"))
		size := internal.ToUint(r.URL.Query().Get("size"))
		limit, offset := internal.PagingToLimitOffset(page, size)

		allReviews, err := reviewRepo.GetAll(limit, offset)
		if err != nil {
			internal.HandleInternalError(w, "Error querying reviewRepo: "+err.Error())
			return
		}

		internal.RespondJson(w, http.StatusOK, allReviews)
	}
}
