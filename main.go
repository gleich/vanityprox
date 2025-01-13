package main

import (
	"net/http"
	"strings"
	"time"

	"github.com/gleich/lumber/v3"
)

func main() {
	setupLogger()
	lumber.Info("booted")

	mux := http.NewServeMux()
	mux.HandleFunc("/", handle)

	server := http.Server{
		Addr:    ":8000",
		Handler: logRequest(mux),
	}
	err := server.ListenAndServe()
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
	if strings.HasSuffix(r.URL.Path, "/info/refs") {
		http.Error(w, "This server does not serve Git repositories.", http.StatusNotFound)
		return
	}

	name := strings.TrimPrefix(r.URL.Path, "/")
	if name == "" {
		http.Redirect(w, r, "https://github.com/gleich/vanityprox", http.StatusMovedPermanently)
		return
	}
	if name == "favicon.ico" {
		http.Error(w, "Not found.", http.StatusNotFound)
		return
	}
	root := strings.Split(name, "/")[0]

	data := templateData{ProjectName: name, ProjectRoot: root}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := htmlTemplate.Execute(w, data)
	if err != nil {
		lumber.Error(err, "failed to execute HTML template")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		lumber.Done(r.Method, r.URL.Path, time.Since(start))
	})
}
