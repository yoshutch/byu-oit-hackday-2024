package main

import (
	"byu.edu/hackday-favorite-color/adapters"
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

	restAdapter, err := adapters.NewRestAdapter(mux)
	if err != nil {
		return
	}
	restAdapter.HandleRoutes()

	log.Print("Listening...")
	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}
