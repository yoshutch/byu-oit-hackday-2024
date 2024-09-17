package pages

import (
	"html/template"
	"log"
	"net/http"
)

type IndexData struct {
	Greeting string
}

func IndexHtmlAdapter(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/index.html")
	if err != nil {
		log.Printf("Error parsing file: %s", err)
		return
	}
	data := IndexData{Greeting: "Hello BYU!"}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %s", err)
		return
	}
}
