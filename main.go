package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func shortenerList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display all shortened URLs..."))
}

func shortenerView(w http.ResponseWriter, r *http.Request) {
	// Get ID and convert to int.
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific shortened URL with ID %d...", id)
}

func shortenerCreate(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Display a form for creating a shortened URL..."))
}

func shortenerCreatePost(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new shortened URL..."))
}

func main() {
	// ServeMux is the route handler.
	mux := http.NewServeMux()

	// Define routes.
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /health", health)
	mux.HandleFunc("GET /shortener", shortenerList)
	mux.HandleFunc("GET /shortener/view/{id}", shortenerView) // Including '/view/...' to avoid overlaps
	mux.HandleFunc("GET /shortener/create", shortenerCreate)
	mux.HandleFunc("POST /shortener/create", shortenerCreatePost)

	// Note log instead of fmt.PrintLn
	log.Print("starting server on :4000")

	// Run server and collect an error if one occurs.
	// Omitted hostname results in listening on all network interfaces.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
