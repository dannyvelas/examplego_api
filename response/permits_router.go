package response

import (
	"github.com/dannyvelas/go-backend/storage"
	"net/http"
)

func ReviewsRouter(reviewRepo storage.ReviewRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.GET("/active", GetActiveReviews(reviewRepo))
	}
}

func GetActiveReviews(reviewRepo storage.ReviewRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		activeReviews, err := reviewRepo.GetActive()

		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}

		respondJson(activeReviews)
	}
}
