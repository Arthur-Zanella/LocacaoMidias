package routes

import (
	"locacaomidias/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRoutes inicializa o router
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/estados", handlers.ListEstados).Methods("GET")
	apiRouter.HandleFunc("/cidades", handlers.ListCidades).Methods("GET")
	apiRouter.HandleFunc("/generos", handlers.ListGeneros).Methods("GET")
	apiRouter.HandleFunc("/tipos", handlers.ListTipos).Methods("GET")
	apiRouter.HandleFunc("/classificacoes-etarias", handlers.ListClassificacoesEtarias).Methods("GET")
	apiRouter.HandleFunc("/classificacoes-internas", handlers.ListClassificacoesInternas).Methods("GET")

	fileServer := http.FileServer(http.Dir("./web"))
	router.PathPrefix("/").Handler(fileServer)

	return router
}
