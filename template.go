package main

import "html/template"

type templateData struct {
	ProjectName string
}

var htmlTemplate = template.Must(template.New("gometa").Parse(`
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8" />
    <meta
      name="go-import"
      content="mattglei.ch/{{.ProjectName}} git https://github.com/gleich/{{.ProjectName}}.git"
    />
    <meta
      name="go-source"
      content="mattglei.ch/{{.ProjectName}} _ https://github.com/gleich/{{.ProjectName}}/tree/main{/dir} https://github.com/gleich/{{.ProjectName}}/blob/main{/dir}/{file}#L{line}"
    />
    <title>mattglei.ch/{{.ProjectName}}</title>
  </head>
  <body>
    <p>
      Redirecting to
      <a href="https://github.com/gleich/{{.ProjectName}}"
        >github.com/gleich/{{.ProjectName}}</a
      >
    </p>
  </body>
</html>
`))
