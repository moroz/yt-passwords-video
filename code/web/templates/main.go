package templates

import (
	"embed"
	"html/template"
)

//go:embed **/*.html.tmpl
var fs embed.FS

func MustParseTemplateWithLayout(path string) *template.Template {
	return template.Must(template.ParseFS(fs, "layout/root.html.tmpl", path))
}

type pageTemplates struct {
	Index *template.Template
}

var Pages = pageTemplates{
	Index: MustParseTemplateWithLayout("pages/index.html.tmpl"),
}
