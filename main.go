package main

import (
	"log"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func shortenerView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display one shortened URL."))
}

func shortenerCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a shortened URL."))
}

func main() {
	// ServeMux is the route handler.
	mux := http.NewServeMux()

	// Define routes.
	mux.HandleFunc("/health", health)
	mux.HandleFunc("/shortener/view", shortenerCreate)
	mux.HandleFunc("/shortener/create", shortenerView)

	// Note log instead of fmt.PrintLn
	log.Print("starting server on :4000")

	// Run server and collect an error if one occurs.
	// Omitted hostname results in listening on all network interfaces.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
