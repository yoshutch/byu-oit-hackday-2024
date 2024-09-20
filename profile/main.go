package main

import (
	"byu.edu/hackday-profile/adapters"
	"byu.edu/hackday-profile/db"
	"byu.edu/hackday-profile/events"
	"byu.edu/hackday-profile/services"
	"log"
	"net/http"
)

func main() {
	// setup dependencies
	profileRepo, err := db.NewProfileRepo("root", "password", 5432, "hackday")
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	adapter, err := events.NewEventAdapter()
	if err != nil {
		log.Fatalf("Error connecting to eventbus: %s", err)
	}
	profileService, err := services.NewProfileService(profileRepo, adapter)
	if err != nil {
		log.Fatalf("Error creating service: %s", err)
	}
	eventAdapter, err := events.NewEventAdapter()
	if err != nil {
		log.Fatalf("Error creating eventAdpater: %s", err)
	}
	log.Printf("here: %s", eventAdapter)

	// http server
	fs := http.FileServer(http.Dir("static/"))
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	htmlAdapter, err := adapters.NewHtmlAdapter(mux, profileService)
	if err != nil {
		log.Fatalf("Failed to instantiate HtmlAdapter: %s", err)
	}
	htmlAdapter.HandleRoutes()

	restAdapter, err := adapters.NewRestAdapter(mux, profileService)
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
