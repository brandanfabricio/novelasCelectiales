package main

import (
	capitulo "Novelas/internal/capitulo/infreastructure/routes"
	novela "Novelas/internal/novela/infrastructure/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	novela.NovelasRoutes(router)
	capitulo.CapituloRoutes(router)

	// router.PathPrefix("/capitulo/").Handler(http.StripPrefix("/cap", routerCapitulo))

	router.PathPrefix("/public/").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))

	// router.Handle("/", routerNovela)
	// rutap.Schemes()

	// router.Handle("/", routerNovela)
	// router.PathPrefix("/").Handler(routerNovela)

	// router.PathPrefix("/").Handler(routerNovela)

	// router.Handle("/capitulo", routerCapitulo)

	log.Println("Server on port 3000")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
