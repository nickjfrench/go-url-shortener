package main

import (
	"context"
	"flag"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
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
	// CLI flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Logger Minimum Level
	logLevel := slog.LevelInfo
	if os.Getenv("DEBUG") == "true" {
		logLevel = slog.LevelDebug
	}

	// Logger Handler
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: logLevel}))

	// Load .env file
	envErr := godotenv.Load()
	if envErr != nil {
		logger.Error("Error loading .env file", slog.Any("error", envErr))
	}

	// DB - using Direct Connection, not a Pool.
	// Note: May need to be changed to Session Pooler and pgxpool
	conn, dbErr := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if dbErr != nil {
		logger.Error("Failed to connect to the database", slog.Any("error", dbErr))
	}
	defer conn.Close(context.Background())

	// Example query to test connection
	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		conn.Close(context.Background())
		logger.Error("Query failed", slog.Any("error", err))
	}
	logger.Info("Successfully connected to DB", slog.String("version", version))

	// Application with DI
	app := &application{
		logger: logger,
	}

	// Using slog.<type> to catch any compile errors.
	logger.Info("starting server", slog.String("addr", *addr))

	// Start http listener using ServeMux route handler.
	if err := http.ListenAndServe(*addr, app.routes()); err != nil {
		logger.Error(err.Error())
		os.Exit(1) // Exit application on error.
	}
}
