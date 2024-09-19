package adapters

import (
	"byu.edu/hackday-profile/db"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type HtmlAdapter struct {
	mux         *http.ServeMux
	indexTmpl   *template.Template
	profileRepo *db.ProfileRepo
}

func NewHtmlAdapter(mux *http.ServeMux, profileRepo *db.ProfileRepo) (*HtmlAdapter, error) {
	// TODO parse all templates into a map of some kind?
	indexTmpl, err := template.ParseFiles("pages/index.html", "pages/layout.html")
	if err != nil {
		log.Printf("Error parsing file: %s", err)
		return nil, err
	}

	return &HtmlAdapter{
		mux,
		indexTmpl,
		profileRepo,
	}, nil
}

func (a HtmlAdapter) HandleRoutes() {
	// Index.html
	a.mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		type IndexData struct {
			Profile *db.Profile
		}
		profile, err := a.profileRepo.GetProfile(1)
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
			fmt.Fprint(w, `<div id="search-results">Error parsing form<div>`)
			return
		}

		fmt.Printf("Request form: %s", r.Form)

		fmt.Fprintf(w, `<div id="results">Saved Successfully!</div>`)
		if err != nil {
			log.Printf("Error executing template: %s", err)
			return
		}
	})
}
