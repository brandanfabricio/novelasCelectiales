package controllers

import (
	"Novelas/internal/capitulo/application"
	"Novelas/internal/web"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var permiseViewPage int = 100

func init() {

	go actualizarVariableDiaria()
}

func actualizarVariableDiaria() {
	for {
		ahora := time.Now()
		// Obtener la hora de inicio del siguiente día
		inicioSiguienteDia := time.Date(ahora.Year(), ahora.Month(), ahora.Day()+1, 0, 0, 0, 0, ahora.Location())
		duracionHastaSiguienteDia := inicioSiguienteDia.Sub(ahora)
		fmt.Println(inicioSiguienteDia)

		<-time.After(duracionHastaSiguienteDia)

		// Actualizar la variable
		permiseViewPage = 1000
		fmt.Println("reset varibl")

	}

}

type CapituloController struct {
	App  application.CaseCapitulointerface
	Tmpl *template.Template
}

func NewCapituloController(capitulo application.CaseCapitulointerface) *CapituloController {

	// Agrega tu función personalizada a las plantillas.
	tmpl := web.NewTemplata()

	// tmpl := template.Must(template.ParseFS(templateFS, "web/layout/*.html", "web/capitulos/*.html"))

	// tmpl := template.Must(template.ParseGlob("web/**/*"))

	return &CapituloController{App: capitulo, Tmpl: tmpl}
}

func (self CapituloController) FindAllCapitulo(w http.ResponseWriter, r *http.Request) {

	user_id := mux.Vars(r)["id"]

	data := self.App.FindAllCapitulo(user_id)

	self.Tmpl.ExecuteTemplate(w, "CapituloIndex", data)

}

type DatosEntrada struct {
	NovelaId string `json:"novelaId"`
	Page     int    `json:"page"`
	// Agrega más campos según sea necesario
}

func (self CapituloController) Paginate(w http.ResponseWriter, r *http.Request) {

	var response DatosEntrada
	err := json.NewDecoder(r.Body).Decode(&response)

	if err != nil {
		http.Error(w, "Error al decodificar los datos JSON de la solicitud", http.StatusBadRequest)
		return
	}

	// user_id := mux.Vars(r)["id"]

	// page := 1

	data := self.App.Paginate(response.NovelaId, response.Page)

	// jsonData,_ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(data)

}

func (this CapituloController) SearchPage(w http.ResponseWriter, r *http.Request) {

	var response DatosEntrada
	err := json.NewDecoder(r.Body).Decode(&response)

	if err != nil {
		http.Error(w, "Error al decodificar los datos JSON de la solicitud", http.StatusBadRequest)
		return
	}

	data := this.App.GetPage(response.NovelaId, response.Page)

	// jsonData,_ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}

type Data struct {
	Id              int
	Titulo          string
	Contenido       template.HTML
	Ncap            int
	NovelaId        string
	Before          int
	Next            int
	PermiseViewPage int
}

func (this CapituloController) FindCapitulo(w http.ResponseWriter, r *http.Request) {

	capituloId, err := strconv.Atoi(r.FormValue("capitulo_id"))
	numeroCap, err := strconv.Atoi(r.FormValue("numero_cap"))
	novelaId := r.FormValue("novela_id")
	if err != nil {
		log.Println("error en el parser de string a numero")
		log.Panicln(err)
	}
	data := this.App.FindCapitulo(capituloId, novelaId, numeroCap)

	before := data.Ncap

	if data.Ncap <= 0 {

		before = 1

	} else {

		before = data.Ncap - 1
	}
	if permiseViewPage > 0 {
		fmt.Println(permiseViewPage)
		permiseViewPage--

	}
	// rebisar
	newData := Data{
		Id:              data.Id,
		Titulo:          data.Titulo,
		Ncap:            data.Ncap,
		NovelaId:        data.NovelaId,
		Contenido:       template.HTML(data.Contenido),
		Before:          before,
		Next:            data.Ncap + 1,
		PermiseViewPage: permiseViewPage,
	}

	// fmt.Printf("titulo %s   antes %d   ahora %d  despues %d ", newData.Titulo, newData.Before, newData.Ncap, newData.Next)

	// newData := "Hello <br/> World"
	this.Tmpl.ExecuteTemplate(w, "readCapitulo", newData)

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(data)

}

type AddCapSt struct {
	Novela_id string `json:"novela_id"`
	Url       string `json:"url"`
}

func (this CapituloController) AddCapitulo(w http.ResponseWriter, r *http.Request) {

	novela_id := r.FormValue("novela_id")
	url := r.FormValue("url")

	this.App.AddCapitulo(novela_id, url)

	urlRedirect := fmt.Sprintf("/capitulo/list/%s", novela_id)

	http.Redirect(w, r, urlRedirect, http.StatusSeeOther)

	// var response AddCapSt
	// err := json.NewDecoder(r.Body).Decode(&response)

	// if err != nil {
	// 	http.Error(w, "Error al decodificar los datos JSON de la solicitud", http.StatusBadRequest)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")

	// json.NewEncoder(w).Encode(response)

	// urlRedirect := fmt.Sprintf("/capitulo/list/%s", url)

	// // query := fmt.Sprintf("SELECT id,titulo,Ncap,novelaId FROM capitulos LIMIT %d OFFSET %d", resultadoPorPagina, offset)

	// http.Redirect(w, r, urlRedirect, 200)

}
