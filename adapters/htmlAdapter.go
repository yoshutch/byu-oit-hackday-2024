package adapters

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type HtmlAdapter struct {
	mux       *http.ServeMux
	indexTmpl *template.Template
}

func NewHtmlAdapter(mux *http.ServeMux) (*HtmlAdapter, error) {
	// TODO parse all templates into a map of some kind?
	indexTmpl, err := template.ParseFiles("pages/index.html", "pages/layout.html")
	if err != nil {
		log.Printf("Error parsing file: %s", err)
		return nil, err
	}
	return &HtmlAdapter{
		mux,
		indexTmpl,
	}, nil
}

func (h HtmlAdapter) HandleRoutes() {
	// Index.html
	h.mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		type IndexData struct {
			PageTitle string
			Greeting  string
		}
		data := IndexData{
			//PageTitle: "Greeting!",
			Greeting: "Hello BYU!",
		}
		err := h.indexTmpl.ExecuteTemplate(w, "layout", data)
		if err != nil {
			log.Printf("Error executing template: %s", err)
			return
		}
	})

	// Ajax calls
	h.mux.HandleFunc("POST /ajax/searchDocs", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Printf("Error parsing form: %s", err)
			fmt.Fprint(w, `<div id="search-results">Error parsing form<div>`)
			return
		}

		fmt.Printf("Request form: %s", r.Form)
		yearTerm := r.FormValue("yearTerm")
		fmt.Fprintf(w, `<div id="search-results">TADA! Got to HTMX for yearTerm: %s!<div>`, yearTerm)
	})
}
