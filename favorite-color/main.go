package main

import (
	"byu.edu/hackday-favorite-color/adapters"
	"byu.edu/hackday-favorite-color/db"
	"byu.edu/hackday-favorite-color/events"
	"byu.edu/hackday-favorite-color/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// env variables
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbDb := os.Getenv("POSTGRES_DB")
	eventHost := os.Getenv("KAFKA_HOST")
	eventPort := os.Getenv("KAFKA_PORT")

	// setup dependencies
	favColorRepo, err := db.NewFavColorRepo(dbHost, dbPort, dbUser, dbPassword, dbDb)
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
	favColorService, err := services.NewFavColorService(favColorRepo)
	if err != nil {
		log.Fatalf("Error creating service: %s", err)
	}

	// http server
	fs := http.FileServer(http.Dir("static/"))
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	htmlAdapter, err := adapters.NewHtmlAdapter(mux, favColorService)
	if err != nil {
		log.Fatalf("Failed to instantiate HtmlAdapter: %s", err)
	}
	htmlAdapter.HandleRoutes()

	restAdapter, err := adapters.NewRestAdapter(mux)
	if err != nil {
		return
	}
	restAdapter.HandleRoutes()

	eventAdapter, err := events.NewEventAdapter(eventHost, eventPort, favColorService)
	if err != nil {
		log.Fatalf("Failed to instantiate event adapter: %s", err)
	}
	//eventAdapter.Close()

	sigChnl := make(chan os.Signal, 1)
	signal.Notify(sigChnl, syscall.SIGINT, syscall.SIGTERM)
	exitchnl := make(chan int)
	go func() {
		for {
			_ = <-sigChnl
			eventAdapter.Close()
		}
	}()

	go serverListen(mux)
	//go func() {
	//	err := eventAdapter.Listen()
	//	if err != nil {
	//		log.Fatalf("Error listening to eventAdapter: %s", err)
	//	}
	//}()
	// wait for exit code
	exitCode := <-exitchnl
	os.Exit(exitCode)
}

func serverListen(mux *http.ServeMux) {
	log.Print("Web Server Listening on :8081 ...")
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatalf("Error listening for web server: %s", err)
	}
}
