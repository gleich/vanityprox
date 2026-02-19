package api

import (
	"net/http"
	"strings"

	"go.mattglei.ch/go.mattglei.ch/internal/conf"
	"go.mattglei.ch/go.mattglei.ch/internal/html"
	"go.mattglei.ch/go.mattglei.ch/internal/pkg"
	"go.mattglei.ch/go.mattglei.ch/internal/util"
)

func handle(config conf.Config, packages *pkg.Packages) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("x-powered-by", "go.mattglei.ch [https://github.com/gleich/go.mattglei.ch]")
		if strings.HasSuffix(r.URL.Path, "/info/refs") {
			http.Error(w, "this server does not serve git repositories", http.StatusNotFound)
			return
		}

		name := strings.TrimPrefix(r.URL.Path, "/")
		if name == "" {
			html.RenderIndex(config, packages, w)
			return
		}
		if name == "favicon.ico" {
			http.Error(w, util.NOT_FOUND_ERROR, http.StatusNotFound)
			return
		}

		html.RenderPackage(config, w, r)
	}
}
