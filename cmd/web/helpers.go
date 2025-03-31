package main

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

// Write a http status to http.ResponseWriter
// Only use in helper functions, unless a helper function for an error type doesn't exist.≈
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Log server error and calls clientError.
// Do call clientError again for the same error, or you'll get `superfluous response` for writing to response twice.
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)

	app.logger.Error(err.Error(), slog.String("method", method), slog.String("uri", uri), slog.String("trace", trace))
	app.clientError(w, 500)
}

func (app *application) logNewRequest(r *http.Request) {
	app.logger.Debug("New Request", slog.Any("method", r.Method), slog.Any("url", r.URL.RequestURI()))
}
