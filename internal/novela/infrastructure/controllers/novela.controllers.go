package controllers

import (
	"Novelas/internal/novela/application"
	"Novelas/internal/web"
	"html/template"
	"net/http"
)

type NovelaController struct {
	App  application.CaseNovelaInt
	Tmpl *template.Template
}

func NewNovelaController(caseNovela application.CaseNovelaInt) *NovelaController {

	// tmpl := lib.LoadTemplate()
	// tmpl := template.Must(template.ParseGlob("./web/**/*.html"))

	// tmpl := template.Must(template.ParseFS(templateFS, "web/**/*.html"))
	tmpl := web.NewTemplata()

	return &NovelaController{App: caseNovela, Tmpl: tmpl} // Tmpl: tmpl,

}

func (app *NovelaController) NovelasAll(w http.ResponseWriter, r *http.Request) {

	data := app.App.FindAllNovela()

	app.Tmpl.ExecuteTemplate(w, "index", data)

}

func (app *NovelaController) SaveNovela(w http.ResponseWriter, r *http.Request) {

	url := r.FormValue("novela-url")
	data := app.App.SaveNovela(url)
	// app.Tmpl.ExecuteTemplate(w, "listNovela", data)
	if data {

		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {

		http.Redirect(w, r, "/", http.StatusNotFound)
	}

}
