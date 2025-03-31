package main

import "net/http"

// Returns a http.ServeMux with our app routes.
func (app *application) routes() *http.ServeMux {
	// Set route handler.
	mux := http.NewServeMux()

	// Set static file server.
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Define routes.
	mux.HandleFunc("GET /health", app.health)
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /url", app.shortenerList)
	mux.HandleFunc("GET /url/view/{id}", app.shortenerView)
	mux.HandleFunc("GET /url/create", app.shortenerCreate)
	mux.HandleFunc("POST /url/create", app.shortenerCreatePost)

	return mux
}
