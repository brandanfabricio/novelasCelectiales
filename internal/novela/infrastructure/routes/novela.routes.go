package routes

import (
	"Novelas/internal/novela/application"
	"Novelas/internal/novela/infrastructure/controllers"
	"Novelas/internal/novela/infrastructure/repository"
	"net/http"

	"github.com/gorilla/mux"
)

func NovelasRoutes(ruter *mux.Router) {

	routerNovela := ruter.PathPrefix("/").Subrouter()

	// config dependecia

	novelaRepo := &repository.MysqlNovelaRepo{}

	novelaCase := application.NewCaseNovela(novelaRepo)

	novelaCtr := controllers.NewNovelaController(novelaCase)

	routerNovela.HandleFunc("/", novelaCtr.NovelasAll).Methods(http.MethodGet)
	routerNovela.HandleFunc("/novela/agregar", novelaCtr.SaveNovela).Methods(http.MethodPost)

}
