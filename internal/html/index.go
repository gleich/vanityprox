package html

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"

	"go.mattglei.ch/vanityprox/internal/conf"
	"go.mattglei.ch/vanityprox/internal/github"
	"go.mattglei.ch/vanityprox/internal/pkg"
	"go.mattglei.ch/vanityprox/internal/util"
)

type Index struct {
	Config   conf.Config
	Packages []github.Repository
}

var (
	//go:embed index.html
	indexHTML     string
	indexTemplate = template.Must(template.New("gometa").Parse(indexHTML))
)

func RenderIndex(config conf.Config, packages *pkg.Packages, w http.ResponseWriter) {
	data := Index{
		Packages: packages.All(),
		Config:   config,
	}
	err := indexTemplate.Execute(w, data)
	if err != nil {
		util.InternalServerError(w, fmt.Errorf(": %w", err))
	}
}
