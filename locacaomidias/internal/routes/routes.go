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

	// atores
	apiRouter.HandleFunc("/atores", handlers.ListAtores).Methods("GET")
	apiRouter.HandleFunc("/atores/{id:[0-9]+}", handlers.GetAtor).Methods("GET")
	apiRouter.HandleFunc("/atores", handlers.CreateAtor).Methods("POST")
	apiRouter.HandleFunc("/atores/{id:[0-9]+}", handlers.UpdateAtor).Methods("PUT")
	apiRouter.HandleFunc("/atores/{id:[0-9]+}", handlers.DeleteAtor).Methods("DELETE")

	// midias
	apiRouter.HandleFunc("/midias", handlers.ListMidias).Methods("GET")
	apiRouter.HandleFunc("/midias/{id:[0-9]+}", handlers.GetMidia).Methods("GET")
	apiRouter.HandleFunc("/midias", handlers.CreateMidia).Methods("POST")
	apiRouter.HandleFunc("/midias/{id:[0-9]+}", handlers.UpdateMidia).Methods("PUT")
	apiRouter.HandleFunc("/midias/{id:[0-9]+}", handlers.DeleteMidia).Methods("DELETE")

	// exemplares
	apiRouter.HandleFunc("/exemplares", handlers.ListExemplares).Methods("GET")
	apiRouter.HandleFunc("/exemplares/{id:[0-9]+}", handlers.GetExemplar).Methods("GET")
	apiRouter.HandleFunc("/exemplares", handlers.CreateExemplar).Methods("POST")
	apiRouter.HandleFunc("/exemplares/{id:[0-9]+}", handlers.UpdateExemplar).Methods("PUT")
	apiRouter.HandleFunc("/exemplares/{id:[0-9]+}", handlers.DeleteExemplar).Methods("DELETE")

	// clientes
	apiRouter.HandleFunc("/clientes", handlers.ListClientes).Methods("GET")
	apiRouter.HandleFunc("/clientes/{id}", handlers.GetCliente).Methods("GET")
	apiRouter.HandleFunc("/clientes", handlers.CreateCliente).Methods("POST")
	apiRouter.HandleFunc("/clientes/{id}", handlers.UpdateCliente).Methods("PUT")
	apiRouter.HandleFunc("/clientes/{id}", handlers.DeleteCliente).Methods("DELETE")

	// locacoes
	apiRouter.HandleFunc("/locacoes", handlers.ListLocacoes).Methods("GET")
	apiRouter.HandleFunc("/locacoes/{id}", handlers.GetLocacao).Methods("GET")
	apiRouter.HandleFunc("/locacoes", handlers.CreateLocacao).Methods("POST")
	apiRouter.HandleFunc("/locacoes/{id}", handlers.UpdateLocacao).Methods("PUT")
	apiRouter.HandleFunc("/locacoes/{id}", handlers.DeleteLocacao).Methods("DELETE")

	return router
}
