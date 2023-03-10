package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/HouseCham/bookings/internal/config"
	"github.com/HouseCham/bookings/internal/models"
	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	//? what am I going to put in the session
	gob.Register(models.Reservation{})

	// set to true when in production
	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour //This session will last for 24 hrs
	session.Cookie.Persist = true     // This cookie will persist after the window is closed
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false // To insist the cookie to be encrypted -> to use only https... in production set to true

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

/* ========== creation of struct myWriter to work as a responseWriter, implementing all methods required ========== */
type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var header http.Header
	return header
}

func (tw *myWriter) WriteHeader(i int) {

}

func (tw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}
