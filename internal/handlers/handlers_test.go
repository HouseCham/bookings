package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"general-q", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"majors-s", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"search-a", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},

	{"post-searchAvailability", "/search-availability", "POST", []postData{
		{key: "start_date", value: "2020-01-01"},
		{key: "end_date", value: "2020-01-02"},
	}, http.StatusOK},
	{"post-searchAvailabilityJSON", "/search-availability-json", "POST", []postData{
		{key: "start_date", value: "2020-01-01"},
		{key: "end_date", value: "2020-01-02"},
	}, http.StatusOK},
	{"post-makeReservation", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Ramses"},
		{key: "last_name", value: "Rmz"},
		{key: "phone", value: "3121332750"},
		{key: "email", value: "cham@gmail.com"},
	}, http.StatusOK},
} 

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	testServer := httptest.NewTLSServer(routes)
	defer testServer.Close()	//? When testServer is not used, this closes it automatically

	for _, e := range theTests {
		if e.method == "GET" {
			resp,err := testServer.Client().Get(testServer.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			data := url.Values{}
			for _, x := range e.params {
				data.Add(x.key, x.value)
			}

			resp, err := testServer.Client().PostForm(testServer.URL + e.url, data)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}
			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s, expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}