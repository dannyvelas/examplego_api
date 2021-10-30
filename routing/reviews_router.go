package routing

import (
	"github.com/dannyvelas/go-backend/storage"
	"github.com/dannyvelas/go-backend/utils"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func ReviewsRouter(reviewRepo storage.ReviewRepo) func(chi.Router) {
	return func(r chi.Router) {
		r.Get("/active", GetActive(reviewRepo))
	}
}

func GetActive(reviewRepo storage.ReviewRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page := utils.ToUint(r.URL.Query().Get("page"))
		size := utils.ToUint(r.URL.Query().Get("size"))
		limit, offset := utils.PagingToLimitOffset(page, size)

		activeReviews, err := reviewRepo.GetActive(limit, offset)
		if err != nil {
			utils.HandleInternalError(w, "Error querying reviewRepo: "+err.Error())
			return
		}

		utils.RespondJson(w, http.StatusOK, activeReviews)
	}
}
