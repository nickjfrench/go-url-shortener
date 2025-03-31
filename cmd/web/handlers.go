package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	app.logNewRequest(r)

	w.Write([]byte("OK"))
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	app.logNewRequest(r)

	// Base must be first
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/partials/footer.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) shortenerList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	app.logNewRequest(r)

	w.Write([]byte("Display all shortened URLs."))
}

func (app *application) shortenerView(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	app.logNewRequest(r)

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		app.serverError(w, r, err)
		app.clientError(w, 404)
		return
	}

	fmt.Fprintf(w, "Display a specific shortened URL with ID %d...", id)
}

func (app *application) shortenerCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	app.logNewRequest(r)

	w.Write([]byte("Display a form for creating a new shortened URL..."))
}

func (app *application) shortenerCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	app.logNewRequest(r)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new shortened URL..."))
}
