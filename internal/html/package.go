package html

import (
	_ "embed"
	"fmt"
	"html/template"
	"net/http"
	"slices"
	"strings"

	"go.mattglei.ch/go.mattglei.ch/internal/conf"
	"go.mattglei.ch/go.mattglei.ch/internal/util"
)

type PackageTemplate struct {
	Name string
	Root string
}

var (
	//go:embed package.html
	packageHTML     string
	packageTemplate = template.Must(template.New("gometa").Parse(packageHTML))
)

func RenderPackage(config conf.Config, w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/")
	root := strings.Split(name, "/")[0]
	if slices.Contains(config.Packages, root) {
		data := PackageTemplate{
			Name: name,
			Root: root,
		}
		err := packageTemplate.Execute(w, data)
		if err != nil {
			util.InternalServerError(w, fmt.Errorf("%w failed to execute HTML template", err))
		}
	} else {
		http.Error(w, util.NOT_FOUND_ERROR, http.StatusNotFound)
	}
}
