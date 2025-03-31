package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

// Application-wide dependency injection
// To access, assign functions to struct. See routes as an example.
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

	// Using slog.<type> to catch any compile errors.
	logger.Info("starting server", slog.String("addr", *addr))

	// Start http listener using ServeMux route handler.
	err := http.ListenAndServe(*addr, app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1) // Exit application on error.
	}
}
