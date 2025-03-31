package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// Application-wide dependencies
type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Logger Minimum Level
	logLevel := slog.LevelInfo
	if os.Getenv("DEBUG") == "true" {
		logLevel = slog.LevelDebug
	}

	// Logger Handler
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: logLevel}))

	// Application with DI
	app := &application{
		logger: logger,
	}

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

	// Using slog.<type> to catch any compile errors.
	logger.Info("starting server", slog.String("addr", *addr))

	// Start http listener using ServeMux route handler.
	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1) // Exit application on error.
	}
}
