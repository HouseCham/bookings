package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/HouseCham/bookings/pkg/config"
	"github.com/HouseCham/bookings/pkg/models"
)

// NewTemplates sets the config for the template package
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(templateData *models.TemplateData) *models.TemplateData{
	return templateData
}

// Render templates using HTML templates
func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {

	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, templateFinded := tc[tmpl]
	if !templateFinded {
		log.Fatal("could not load template from cache")
	}

	buf := new(bytes.Buffer)

	templateData = AddDefaultData(templateData)
	_ = t.Execute(buf, templateData)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages {

		// foreach template, create new object
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// look for layout.html files
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		// merge the layout file with the template
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		// finally add the template to the map
		myCache[name] = ts
	}

	return myCache, nil
}
