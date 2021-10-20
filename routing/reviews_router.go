package routing

import (
	"fmt"
	"net/http"

	"github.com/dannyvelas/go-backend/storage"
	"github.com/dannyvelas/go-backend/utils"
	"github.com/go-chi/chi/v5"
)

func ReviewsRouter(reviewRepo storage.ReviewRepo) func(chi.Router) {
	return func(r chi.Router) {
		r.Get("/active", GetActiveReviews(reviewRepo))
	}
}

func GetActiveReviews(reviewRepo storage.ReviewRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page := utils.ToUint(r.URL.Query().Get("page"))
		size := utils.ToUint(r.URL.Query().Get("size"))

		limit, offset := utils.PagingToLimitOffset(page, size)
		activeReviews, err := reviewRepo.GetActive(limit, offset)

		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}

		utils.RespondJson(w, http.StatusOK, activeReviews)
	}
}
