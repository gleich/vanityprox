package main

import "html/template"

type templateData struct {
	ProjectName string
	ProjectRoot string
}

var htmlTemplate = template.Must(template.New("gometa").Parse(`<!DOCTYPE html>
<html>
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width" />
    <meta
      name="go-import"
      content="pkg.mattglei.ch/{{.ProjectName}} git https://github.com/gleich/{{.ProjectRoot}}"
    />
    <meta
      name="go-source"
      content="pkg.mattglei.ch/{{.ProjectName}} https://github.com/gleich/{{.ProjectRoot}} https://github.com/gleich/{{.ProjectRoot}}/tree/main/{/dir} https://github.com/gleich/{{.ProjectRoot}}/blob/main{/dir}/{file}#L{line}"
    />
    <meta http-equiv="refresh" content="0; url=https://pkg.go.dev/github.com/gleich/{{.ProjectName}}">
    <link rel="icon" href="https://mattglei.ch/favicon.ico">
    <title>{{.ProjectName}}</title>
  </head>

   <body>
    <a href="https://github.com/gleich/{{.ProjectRoot}}">github.com/gleich/{{.ProjectRoot}}</a><br>
    <a href="https://github.com/gleich/vanityprox" target="_blank">Proxied by gleich/vanityprox</a>
  </body>
</html>
`))
