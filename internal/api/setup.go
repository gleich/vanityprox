package api

import (
	"net/http"

	"go.mattglei.ch/go.mattglei.ch/internal/conf"
	"go.mattglei.ch/go.mattglei.ch/internal/github"
	"go.mattglei.ch/go.mattglei.ch/internal/pkg"
)

func Setup(config conf.Config, clients github.Clients, packages *pkg.Packages) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", rootEndpoint(config, packages))
	mux.HandleFunc("GET /health", healthEndpoint)
	mux.HandleFunc("POST /github/webhook", func(w http.ResponseWriter, r *http.Request) {
		webhookEndpoint(w, r, clients, packages)
	})
	mux.HandleFunc("GET /styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "styles.css")
	})

	server := http.Server{
		Addr:    ":8000",
		Handler: logRequest(mux),
	}
	return &server
}
