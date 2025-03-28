package main

import (
	"log"
	"net/http"
)

func main() {
	// Set route handler.
	mux := http.NewServeMux()

	// Set static file server.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /health", health)
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /url", shortenerList)
	mux.HandleFunc("GET /url/view/{id}", shortenerView)
	mux.HandleFunc("GET /url/create", shortenerCreate)
	mux.HandleFunc("POST /url/create", shortenerCreatePost)

	port := ":4000"

	log.Printf("starting server on %s", port)

	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
