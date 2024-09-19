package main

import (
	"byu.edu/hackday-favorite-color/adapters"
	"byu.edu/hackday-favorite-color/events"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	eventAdapter, err := events.NewEventAdapter()
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
	go func() {
		err := eventAdapter.Listen()
		if err != nil {
			log.Fatalf("Error listening to eventAdapter: %s", err)
		}
	}()
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
