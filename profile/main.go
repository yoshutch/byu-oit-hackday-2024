package main

import (
	"byu.edu/hackday-profile/adapters"
	"byu.edu/hackday-profile/db"
	"log"
	"net/http"
)

func main() {
	profileRepo, err := db.NewProfileRepo("root", "password", 5433, "hackday")
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	mux := http.NewServeMux()

	htmlAdapter, err := adapters.NewHtmlAdapter(mux, profileRepo)
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
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
