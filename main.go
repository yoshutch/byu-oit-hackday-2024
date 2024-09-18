package main

import (
	"byu.edu/poc-imaging/pages"
	"log"
	"net/http"
)

func main() {
	//mux := http.NewServeMux()

	http.HandleFunc("GET /index", pages.IndexHtmlAdapter)

	log.Print("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
