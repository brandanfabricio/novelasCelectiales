package routes

import (
	"Novelas/internal/user/app"
	"Novelas/internal/user/infreastructure/controllers"
	"Novelas/internal/user/infreastructure/repository"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRoute(router *mux.Router) {

	routerUser := router.PathPrefix("/user").Subrouter()

	// config dependecia
	repoUser := &repository.UserMysql{}
	caseUser := app.NewCaseUser(repoUser)
	userCtr := controllers.NewUserControlers(caseUser)

	// rutas
	routerUser.HandleFunc("/login", userCtr.LoginTemplate).Methods(http.MethodGet)
	routerUser.HandleFunc("/login", userCtr.Login).Methods(http.MethodPost)

	routerUser.HandleFunc("/signUp", userCtr.SignUp).Methods(http.MethodGet)
	routerUser.HandleFunc("/signUp", userCtr.SaveUser).Methods(http.MethodPost)
	routerUser.HandleFunc("/{id}", userCtr.FindUser).Methods(http.MethodGet)

}
