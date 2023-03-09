package handlers

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"html/template"
	"time"

	"github.com/HouseCham/bookings/internal/config"
	"github.com/HouseCham/bookings/internal/models"
	"github.com/HouseCham/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/justinas/nosurf"
)

var app config.AppConfig
var session *scs.SessionManager
var functions = template.FuncMap{}

var pathToTemplates string = "./../../templates"

func getRoutes() http.Handler {
	//? what am I going to put in the session
	gob.Register(models.Reservation{})

	// set to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour //This session will last for 24 hrs
	session.Cookie.Persist = true     // This cookie will persist after the window is closed
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // To insist the cookie to be encrypted -> to use only https... in production set to true

	app.Session = session

	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = true

	repo := NewRepo(&app)

	// Share config file with handlers.go
	NewHandlers(repo)

	// Share config file with render.go
	render.NewTemplates(&app)

	//! /* ======================================================= */
	
	mux := chi.NewRouter()

	// CHI middlewares
	mux.Use(middleware.Recoverer)
	
	// using my own middlewares
	mux.Use(NoSurf)	//? Ignore anny request POST that doesn't have a propper Cross side forgery token
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/generals-quarters", Repo.Generals)
	mux.Get("/majors-suite", Repo.Majors)
	mux.Get("/search-availability", Repo.Availability)
	mux.Get("/contact", Repo.Contact)
	mux.Get("/make-reservation", Repo.MakeReservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)
	
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Post("/search-availability-json", Repo.JSONPostAvailability)
	mux.Post("/make-reservation", Repo.PostReservation)

	//? File server -> in order to pull files from the project like images, css & js
	fileServer := http.FileServer(http.Dir("./src/"))
	mux.Handle("/src/*", http.StripPrefix("/src", fileServer))

	return mux
}

// adds CSRF protection to all POST request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

func CreateTestTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.html from ./templates
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages {

		// foreach template, create new object
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// look for layout.html files
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		// merge the layout file with the template
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.html", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		// finally add the template to the map
		myCache[name] = ts
	}

	return myCache, nil
}