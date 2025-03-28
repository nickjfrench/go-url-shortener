package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	w.Write([]byte("OK"))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	// Base must be first
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/partials/footer.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func shortenerList(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	w.Write([]byte("Display all shortened URLs."))
}

func shortenerView(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific shortened URL with ID %d...", id)
}

func shortenerCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	w.Write([]byte("Display a form for creating a new shortened URL..."))
}

func shortenerCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new shortened URL..."))
}
