package main

import (
	"net/http"
	"strings"
	"time"

	"go.mattglei.ch/timber"
)

const port = ":8000"

func main() {
	conf, err := readConfig()
	if err != nil {
		timber.Fatal(err, "failed to load configuration")
	}

	setupLogger(conf)
	conf.log()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handle(conf))

	mux.HandleFunc("GET /styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "styles.css")
	})

	server := http.Server{
		Addr:    port,
		Handler: logRequest(mux),
	}

	timber.Donef("starting server on 0.0.0.0%s", port)
	err = server.ListenAndServe()
	if err != nil {
		timber.Fatal(err, "failed to start server")
	}
}

func setupLogger(conf config) {
	timezone := conf.Logs.Timezone
	if timezone != "" {
		location, err := time.LoadLocation(timezone)
		if err != nil {
			timber.Fatal(err, "failed to load timezone:", timezone)
		}
		timber.Timezone(location)
	}
	format := conf.Logs.TimeFormat
	if format != "" {
		timber.TimeFormat(format)
	}
}

const NOT_FOUND_ERROR = "package not found"

func handle(conf config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("x-powered-by", "vanityprox [https://github.com/gleich/vanityprox]")
		if strings.HasSuffix(r.URL.Path, "/info/refs") {
			http.Error(w, "this server does not serve git repositories", http.StatusNotFound)
			return
		}

		name := strings.TrimPrefix(r.URL.Path, "/")
		if name == "" {
			renderIndexHTML(conf, w)
			return
		}
		if name == "favicon.ico" {
			http.Error(w, NOT_FOUND_ERROR, http.StatusNotFound)
			return
		}

		renderPackageHTML(conf, w, r)
	}
}

func internalServerError(w http.ResponseWriter, err error) {
	timber.Error(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
