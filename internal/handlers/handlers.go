package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/HouseCham/bookings/internal/config"
	"github.com/HouseCham/bookings/internal/forms"
	"github.com/HouseCham/bookings/internal/models"
	"github.com/HouseCham/bookings/internal/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

//? NewRepo creates a new repository -> Gets repo from main.go
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIP", remoteIP)

	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

// About page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(r.Context(), "remoteIP")
	stringMap["remoteIP"] = remoteIP

	// send data to the template
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}

func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{})
}

func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start_date")
	end := r.Form.Get("end_date")

	w.Write([]byte(fmt.Sprintf("Start date is %s while end date is %s", start, end)))
}

type jsonResponse struct {
	Ok      bool `json:"ok"`
	Message string `json:"message"`
	ArriveDate string `json:"arriveDate"`
	DepartDate string `json:"departDate"`
}

func (m *Repository) JSONPostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start_date")
	end := r.Form.Get("end_date")

	response := jsonResponse {
		Ok: true,
		Message: "Available!!",
		ArriveDate: start,
		DepartDate: end,
	}

	out, err := json.MarshalIndent(response, "", "     ")
	if err != nil {
		log.Println(err)
	}

	log.Println(string(out))

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	//w.Write([]byte(fmt.Sprintf("Start date is %s while end date is %s", start, end)))
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}

func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "makeReservation.page.html", &models.TemplateData{
		Form: forms.New(nil),
	})
}

// Handles the posting of a reservation form
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	
}