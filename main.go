package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"pkg.mattglei.ch/timber"
)

func main() {
	conf, err := readConfig()
	if err != nil {
		timber.Fatal(err, "failed to load configuration")
	}

	setupLogger(conf)
	logConfig(conf)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handle(conf))

	server := http.Server{
		Addr:    ":8000",
		Handler: logRequest(mux),
	}

	timber.Done("starting server")
	err = server.ListenAndServe()
	if err != nil {
		timber.Fatal(err, "failed to start server")
	}
}

func setupLogger(conf config) {
	if conf.LogTimezone != "" {
		timezone, err := time.LoadLocation(conf.LogTimezone)
		if err != nil {
			timber.Fatal(err, "failed to load timezone:", conf.LogTimezone)
		}
		timber.SetTimezone(timezone)
	}
	if conf.LogTimeFormat != "" {
		timber.SetTimeFormat(conf.LogTimeFormat)
	}
}

func handle(conf config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/info/refs") {
			http.Error(w, "This server does not serve Git repositories.", http.StatusNotFound)
			return
		}

		name := strings.TrimPrefix(r.URL.Path, "/")
		if name == "" {
			if conf.RootRedirect != "" {
				http.Redirect(w, r, conf.RootRedirect, http.StatusMovedPermanently)
				return
			} else {
				http.Error(w, "Not found.", http.StatusNotFound)
				return
			}
		}
		if name == "favicon.ico" {
			http.Error(w, "Not found.", http.StatusNotFound)
			return
		}
		root := strings.Split(name, "/")[0]

		w.Header().Set("content-type", "text/html; charset=utf-8")
		w.Header().Set("x-powered-by", "vanityprox")

		data := templateData{ProjectName: name, ProjectRoot: root, Config: conf}
		err := htmlTemplate.Execute(w, data)
		if err != nil {
			err = fmt.Errorf("%v failed to execute HTML template", err)
			timber.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		timber.Done(r.URL.Path, time.Since(start))
	})
}
