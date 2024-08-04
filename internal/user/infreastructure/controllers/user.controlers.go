package controllers

import (
	"Novelas/internal/user/domain"
	"Novelas/internal/web"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserControlers struct {
	// App  application.CaseNovelaInt
	App  domain.CaseUserInt
	Tmpl *template.Template
}

func NewUserControlers(caseUser domain.CaseUserInt) *UserControlers {
	tmpl := web.NewTemplata()
	return &UserControlers{App: caseUser, Tmpl: tmpl} // Tmpl: tmpl,
}

func (app *UserControlers) FindUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Error al convertir:", err)
		return
	}
	data := app.App.FindUser(id)
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (app *UserControlers) SaveUser(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	newUser := domain.User{
		Id:       "",
		Name:     name,
		Email:    email,
		Password: password,
		Status:   "activo",
		RolId:    "",
	}
	response, err := app.App.SaveUser(newUser)
	if err != nil {
		app.Tmpl.ExecuteTemplate(w, "signUp", response)
	} else {
		fmt.Println(response)
		if response != 201 {
			app.Tmpl.ExecuteTemplate(w, "signUp", response)

		} else {
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)

		}
	}

	// var response domain.User
	// err := json.NewDecoder(r.Body).Decode(&response)
	// if err != nil {
	// 	http.Error(w, "Error al decodificar los datos JSON de la solicitud", http.StatusBadRequest)
	// 	return
	// }
	// app.App.SaveUser(response)
}

func (app *UserControlers) SignUp(w http.ResponseWriter, r *http.Request) {

	app.Tmpl.ExecuteTemplate(w, "signUp", nil)

}
func (app *UserControlers) LoginTemplate(w http.ResponseWriter, r *http.Request) {

	app.Tmpl.ExecuteTemplate(w, "login", nil)

}
func (app *UserControlers) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	newUser := domain.User{
		Id:       "",
		Name:     "",
		Email:    email,
		Password: password,
		Status:   "activo",
		RolId:    "",
	}
	response, err := app.App.Login(newUser)

	fmt.Println(response)

	if err != nil {
		app.Tmpl.ExecuteTemplate(w, "login", response)
	} else {
		if response != 201 {
			app.Tmpl.ExecuteTemplate(w, "login", response)

		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)

		}
	}
	app.Tmpl.ExecuteTemplate(w, "login", nil)

}
