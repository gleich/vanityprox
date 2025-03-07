package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"go.mattglei.ch/timber"
)

var (
	cache      map[string][]byte = map[string][]byte{}
	cacheMutex sync.RWMutex
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

	timber.Done("starting server on 0.0.0.0:8000")
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
		timber.Timezone(timezone)
	}
	if conf.LogTimeFormat != "" {
		timber.TimeFormat(conf.LogTimeFormat)
	}
}

const NOT_FOUND_ERROR = "requested resource not found"

func handle(conf config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("x-powered-by", "vanityprox [https://github.com/gleich/vanityprox]")
		if strings.HasSuffix(r.URL.Path, "/info/refs") {
			http.Error(w, "this server does not serve git repositories", http.StatusNotFound)
			return
		}

		name := strings.TrimPrefix(r.URL.Path, "/")
		if name == "" {
			if conf.RootRedirect != "" {
				http.Redirect(w, r, conf.RootRedirect, http.StatusMovedPermanently)
				return
			} else {
				http.Error(w, NOT_FOUND_ERROR, http.StatusNotFound)
				return
			}
		}
		if name == "favicon.ico" {
			http.Error(w, NOT_FOUND_ERROR, http.StatusNotFound)
			return
		}

		cacheMutex.RLock()
		cachedTemplate, found := cache[name]
		cacheMutex.RUnlock()
		if found {
			_, err := w.Write(cachedTemplate)
			if err != nil {
				err = fmt.Errorf("%w failed to write cached HTML template", err)
				timber.Error(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		root := strings.Split(name, "/")[0]
		data := templateData{ProjectName: name, ProjectRoot: root, Config: conf}
		var buf bytes.Buffer
		err := htmlTemplate.Execute(&buf, data)
		if err != nil {
			err = fmt.Errorf("%w failed to execute HTML template", err)
			timber.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result := buf.Bytes()

		cacheMutex.Lock()
		cache[name] = result
		cacheMutex.Unlock()

		_, err = w.Write(result)
		if err != nil {
			err = fmt.Errorf("%w failed to write new html template to response", err)
			timber.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
