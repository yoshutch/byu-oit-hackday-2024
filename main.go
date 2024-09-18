package main

import (
	"byu.edu/poc-imaging/adapters"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	htmlAdapter, err := adapters.NewHtmlAdapter(mux)
	if err != nil {
		log.Fatalf("Failed to instantiate HtmlAdapter: %s", err)
	}
	htmlAdapter.HandleRoutes()

	log.Print("Listening...")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
