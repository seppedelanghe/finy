package rendering

import (
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/unrolled/render"
)


var Renderer = render.New(render.Options{
    Directory: "templates", // Specify what path to load the templates from.
    Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
	Funcs: []template.FuncMap{
		{ "IncludeHTML": IncludeHTML, },
		{ "IncludeIcon": IncludeIcon, },
	},
})


func IncludeHTML(path string) template.HTML {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return template.HTML(string(b))
}

func IncludeIcon(name string) template.HTML {
	path := filepath.Join("assets/icons/", name)
	b, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
		return ""
	}

	return template.HTML(string(b))
}
