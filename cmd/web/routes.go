package main

import (
	"net/http"

	//"github.com/bmizerany/pat"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/middleware"
	"github.com/HouseCham/bookings/pkg/config"
	"github.com/HouseCham/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	// CHI middlewares
	mux.Use(middleware.Recoverer)
	
	// using my own middlewares
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
