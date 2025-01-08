package main

import (
	"net/http"
	"time"

	"github.com/gleich/lumber/v3"
)

func main() {
	setupLogger()
	lumber.Info("booted")

	mux := http.NewServeMux()
	mux.HandleFunc("/{name}", handle)

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		lumber.Fatal(err, "failed to start server")
	}
}

func setupLogger() {
	nytime, err := time.LoadLocation("America/New_York")
	if err != nil {
		lumber.Fatal(err, "failed to load new york timezone")
	}
	lumber.SetTimezone(nytime)
	lumber.SetTimeFormat("01/02 03:04:05 PM MST")
}

func handle(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	if name == "" {
		http.Error(w, "Project name not specified", http.StatusNotFound)
		return
	}

	data := templateData{ProjectName: name}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := htmlTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
