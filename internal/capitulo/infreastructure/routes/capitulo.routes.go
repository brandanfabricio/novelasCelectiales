package routes

import (
	"Novelas/internal/capitulo/application"
	"Novelas/internal/capitulo/infreastructure/controllers"
	capitulo "Novelas/internal/capitulo/infreastructure/repository"
	novela "Novelas/internal/novela/infrastructure/repository"
	"net/http"

	"github.com/gorilla/mux"
)

func CapituloRoutes(ruter *mux.Router) {

	routerCapitulo := ruter.PathPrefix("/capitulo").Subrouter()

	// config dependecia

	// repo
	novelaRepo := &novela.MysqlNovelaRepo{}
	capituloRepo := &capitulo.MysqlCApituloRepo{}

	// casos
	capituloCase := application.NewCaseCApitulo(novelaRepo, capituloRepo)

	// controller

	capituloCtr := controllers.NewCapituloController(capituloCase)

	// router.HandleFunc("/", novelaCtr.NovelasAll).Methods(http.MethodGet)
	// router.HandleFunc("/novela/agregar", novelaCtr.SaveNovela).Methods(http.MethodPost)

	routerCapitulo.HandleFunc("/list/{id}", capituloCtr.FindAllCapitulo).Methods(http.MethodGet)
	routerCapitulo.HandleFunc("/list/paginate", capituloCtr.Paginate).Methods(http.MethodPost)
	routerCapitulo.HandleFunc("/list/searchPage", capituloCtr.SearchPage).Methods(http.MethodPost)
	routerCapitulo.HandleFunc("/list/read", capituloCtr.FindCapitulo).Methods(http.MethodPost)

	routerCapitulo.HandleFunc("/add", capituloCtr.AddCapitulo).Methods(http.MethodPost)

}
