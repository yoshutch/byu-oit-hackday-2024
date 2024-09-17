package main

import (
	"byu.edu/poc-imaging/pages"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", pages.IndexHtmlAdapter)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
