package adapters

import (
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
	indexTmpl, err := template.ParseFiles("pages/index.html")
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
			Greeting string
		}
		data := IndexData{Greeting: "Hello BYU!"}
		err := h.indexTmpl.Execute(w, data)
		if err != nil {
			log.Printf("Error executing template: %s", err)
			return
		}
	})
}
