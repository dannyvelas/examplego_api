package main

import (
	"github.com/dannyvelas/go-backend/config"
	"github.com/dannyvelas/go-backend/routing"
	"github.com/dannyvelas/go-backend/storage"
	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
)

func main() {
	config, err := config.New()
	if err != nil {
		panic(err)
	}

	database, err := storage.NewDatabase(config.Postgres)
	if err != nil {
		panic(err)
	}

	adminRepo := storage.NewAdminRepo(database)
	reviewRepo := storage.NewReviewRepo(database)

	router := chi.NewRouter()
	router.Post("/api/login", routing.HandleLogin(*adminRepo))
	router.Route("/api/reviews", routing.ReviewsRouter(*reviewRepo))

	log.Fatal(http.ListenAndServe(":5000", router))
}
