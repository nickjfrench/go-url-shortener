package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

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

	log.Printf("starting server on %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
