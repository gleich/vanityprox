package api

import (
	"net/http"

	"go.mattglei.ch/vanityprox/internal/conf"
	"go.mattglei.ch/vanityprox/internal/pkg"
)

func Setup(config conf.Config, packages *pkg.Packages) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handle(config, packages))

	mux.HandleFunc("GET /styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "styles.css")
	})

	server := http.Server{
		Addr:    ":8000",
		Handler: logRequest(mux),
	}

	return &server
}
