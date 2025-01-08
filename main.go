package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gleich/lumber/v3"
)

func main() {
	setupLogger()
	lumber.Info("booted")

	mux := http.NewServeMux()
	mux.HandleFunc("/", handle)

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
	name := strings.TrimPrefix(r.URL.Path, "/")
	if name == "" {
		http.Error(w, "Project name not specified", http.StatusNotFound)
		return
	}

	redirectURL, err := url.JoinPath("https://github.com/gleich/", name)
	if err != nil {
		err = fmt.Errorf("%v failed to join path", err)
		lumber.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
	lumber.Done("redirected", name, "->", redirectURL)
}
