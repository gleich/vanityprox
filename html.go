package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"slices"
	"strings"
)

type packageData struct {
	ProjectName string
	ProjectRoot string
	Config      config
}

//go:embed package.html
var packageHTML string

var packageTemplate = template.Must(template.New("gometa").Parse(packageHTML))

func renderPackageHTML(conf config, w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/")
	root := strings.Split(name, "/")[0]
	if slices.Contains(conf.Packages, root) {
		data := packageData{
			ProjectName: name,
			ProjectRoot: root,
			Config:      conf,
		}
		err := packageTemplate.Execute(w, data)
		if err != nil {
			internalServerError(w, fmt.Errorf("%w failed to execute HTML template", err))
		}
	} else {
		http.Error(w, NOT_FOUND_ERROR, http.StatusNotFound)
	}
}

type indexData struct {
	Packages []string
	Config   config
}

//go:embed index.html
var indexHTML string

var indexTemplate = template.Must(template.New("gometa").Parse(indexHTML))

func renderIndexHTML(conf config, w http.ResponseWriter) {
	data := indexData{Packages: []string{"hello"}, Config: conf}
	err := indexTemplate.Execute(w, data)
	if err != nil {
		internalServerError(w, fmt.Errorf(": %w", err))
	}
}
