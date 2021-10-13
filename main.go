package main

import (
	"fmt"
	"github.com/dannyvelas/go-backend/config"
	"github.com/dannyvelas/go-backend/response"
	"github.com/dannyvelas/go-backend/storage"
	_ "github.com/joho/godotenv/autoload"
	"github.com/julienschmidt/httprouter"
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

	reviewRepo := storage.NewReviewRepo(database)

	router := httprouter.New()
	router.HandlerFunc("/api/reviews", response.ReviewsRouter(reviewRepo))

	log.Fatal(http.ListenAndServe(":5000", router))
}
