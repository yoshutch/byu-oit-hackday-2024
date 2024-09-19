package adapters

import (
	"byu.edu/hackday-profile/db"
	"byu.edu/hackday-profile/services"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type HtmlAdapter struct {
	mux            *http.ServeMux
	indexTmpl      *template.Template
	profileService *services.ProfileService
}

func NewHtmlAdapter(mux *http.ServeMux, profileService *services.ProfileService) (*HtmlAdapter, error) {
	// TODO parse all templates into a map of some kind?
	indexTmpl, err := template.ParseFiles("pages/index.html", "pages/layout.html")
	if err != nil {
		log.Printf("Error parsing file: %s", err)
		return nil, err
	}

	return &HtmlAdapter{
		mux,
		indexTmpl,
		profileService,
	}, nil
}

func (a HtmlAdapter) HandleRoutes() {
	// Index.html
	a.mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		type IndexData struct {
			Profile *db.Profile
		}
		profile, err := a.profileService.LoadProfile(1)
		if err != nil {
			log.Printf("Error getting profile: %s", err)
			return
		}

		data := IndexData{
			//PageTitle: "Greeting!",
			Profile: profile,
		}
		err = a.indexTmpl.ExecuteTemplate(w, "layout", data)
		if err != nil {
			log.Printf("Error executing template: %s", err)
			return
		}
	})

	// Ajax calls
	a.mux.HandleFunc("POST /ajax/saveProfile", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Printf("Error parsing form: %s", err)
			fmt.Fprint(w, `<div id="results">Error parsing form</div>`)
			return
		}

		fmt.Printf("Request form: %s", r.Form)
		err = a.profileService.SaveProfile(1, r.FormValue("FirstName"), r.FormValue("LastName"))
		if err != nil {
			log.Printf("Error saving profile: %s", err)
			fmt.Fprint(w, `<div id="results">Error saving profile</div>`)
			return
		}

		fmt.Fprintf(w, `<div id="results">Saved Successfully!</div>`)
		if err != nil {
			log.Printf("Error executing template: %s", err)
			return
		}
	})
}
