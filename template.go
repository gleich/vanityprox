package main

import (
	_ "embed"
	"html/template"
)

type templateData struct {
	ProjectName string
	ProjectRoot string
	Config      config
}

//go:embed template.html
var html string

var htmlTemplate = template.Must(template.New("gometa").Parse(html))
