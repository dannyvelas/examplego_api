package main

import (
	"context"
	"fmt"
	"github.com/dannyvelas/examplego_api/config"
	"github.com/dannyvelas/examplego_api/routing"
	"github.com/dannyvelas/examplego_api/storage"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Info().Msg("Initializing app...")

	// load config
	config := config.NewConfig()

	// connect to database
	// no defer close() because connection closes automatically on program exit
	database, err := storage.NewDatabase(config.Postgres())
	if err != nil {
		log.Fatal().Msgf("Failed to start database: %s", err)
		return
	}
	log.Info().Msg("Connected to Database.")

	// initialize repos
	adminRepo := storage.NewAdminRepo(database)
	reviewRepo := storage.NewReviewRepo(database)

	// initialize JWTMiddleware
	jwtMiddleware := routing.NewJWTMiddleware(config.Token())

	// set routes
	router := chi.NewRouter()
	router.Route("/api", func(apiRouter chi.Router) {
		apiRouter.Post("/login", routing.Login(jwtMiddleware, adminRepo))
		apiRouter.Route("/admin", func(adminRouter chi.Router) {
			adminRouter.Use(jwtMiddleware.Authenticate)
			adminRouter.Route("/hello", routing.HelloRouter())
			adminRouter.Route("/reviews", routing.ReviewsRouter(reviewRepo))
		})
	})

	// configure http server
	httpConfig := config.Http()
	httpServer := http.Server{
		Addr:         fmt.Sprintf("%s:%d", httpConfig.Host(), httpConfig.Port()),
		Handler:      router,
		ReadTimeout:  httpConfig.ReadTimeout(),
		WriteTimeout: httpConfig.WriteTimeout(),
		IdleTimeout:  httpConfig.IdleTimeout(),
	}

	// initialize error channel
	errChannel := make(chan error)
	defer close(errChannel)

	// receive errors from startup or signal interrupt
	go func() {
		errChannel <- StartServer(httpServer)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChannel <- fmt.Errorf("%s", <-c)
	}()

	fatalErr := <-errChannel
	log.Info().Msgf("Closing server: %s", fatalErr)

	shutdownGracefully(30*time.Second, httpServer)
}

func StartServer(httpServer http.Server) error {
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal().Msgf("Failed to start server: %s", err)
		return err
	}
	return nil
}

func shutdownGracefully(timeout time.Duration, httpServer http.Server) {
	log.Info().Msg("Gracefully shutting down...")

	gracefullCtx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := httpServer.Shutdown(gracefullCtx); err != nil {
		log.Error().Msgf("Error shutting down the server: %s", err)
	} else {
		log.Info().Msg("HttpServer gracefully shut down")
	}
}