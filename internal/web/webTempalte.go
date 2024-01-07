package web

import (
	"embed"
	"html/template"
)

//go:embed layout/*.html capitulos/*.html novela/*.html
var templateFS embed.FS

func NewTemplata() *template.Template {

	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}

	// Agrega tu función personalizada a las plantillas.
	tmpl := template.Must(template.New("").Funcs(funcMap).ParseFS(templateFS, "layout/*.html", "capitulos/*.html","novela/*.html"))

	return tmpl
}
