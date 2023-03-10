package render

import (
	"github.com/HouseCham/bookings/internal/models"
	"net/http"
	"testing"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	response, err := getSession()
	if err != nil {
		t.Error(err)
	}

	// just adding "123" as a variable session
	session.Put(response.Context(), "flash", "123")
	result := AddDefaultData(&td, response)
	if result.Flash != "123" {
		t.Error("flash value 123 not found in the session")
	}
}

func getSession() (*http.Request, error) {
	response, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	// Getting and returning session data
	context := response.Context()
	context, _ = session.Load(context, response.Header.Get("X-Session"))
	response = response.WithContext(context)

	return response, nil
}
