package main

import (
	"net/http"

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
	mux.Use(NoSurf)	// Ignore anny request POST that doesn't have a propper Cross side forgery token
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)
	
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.JSONPostAvailability)

	// File server -> in order to pull files from the project like images, css & js
	fileServer := http.FileServer(http.Dir("./src/"))
	mux.Handle("/src/*", http.StripPrefix("/src", fileServer))

	return mux
}
