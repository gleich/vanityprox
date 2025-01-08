package main

import "html/template"

type templateData struct {
	ProjectName string
	ProjectRoot string
}

var htmlTemplate = template.Must(template.New("gometa").Parse(`
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta
      name="go-import"
      content="pkg.mattglei.ch/{{.ProjectName}} git https://github.com/gleich/{{.ProjectRoot}}.git"
    />
    <meta http-equiv="refresh" content="0; url=https://pkg.go.dev/{{.ProjectName}}">
    <title>{{.ProjectName}}</title>
  </head>
</html>
`))
