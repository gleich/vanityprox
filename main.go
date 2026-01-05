package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"

	"go.mattglei.ch/timber"
)

const port = ":8000"

var client = http.Client{
	Timeout: 10 * time.Second,
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	},
}

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
				http.Redirect(w, r, conf.RootRedirect, http.StatusPermanentRedirect)
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

		// check to make sure that requested resource actually exists
		root := strings.Split(name, "/")[0]
		repoURL := fmt.Sprintf("https://%s/%s", conf.SourcePrefix, root)
		resp, err := client.Head(repoURL)
		if err != nil {
			err = fmt.Errorf("%w failed to make HEAD request to %s", err, repoURL)
			timber.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Cache-Control", "public, max-age=3600")
		if resp.StatusCode == http.StatusOK {
			data := templateData{ProjectName: name, ProjectRoot: root, Config: conf}
			var buf bytes.Buffer
			err = htmlTemplate.Execute(&buf, data)
			if err != nil {
				err = fmt.Errorf("%w failed to execute HTML template", err)
				timber.Error(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			_, err = w.Write(buf.Bytes())
			if err != nil {
				err = fmt.Errorf("%w failed to write new html template to response", err)
				timber.Error(err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		} else {
			http.Error(w, resp.Status, resp.StatusCode)
		}
		w.WriteHeader(http.StatusOK)
	}
}
