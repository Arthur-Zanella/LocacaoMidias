package routes

import (
	"locacaomidias/internal/handlers"

	"github.com/gorilla/mux"
)

// InitRoutes inicializa o router da aplicação
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/estados", handlers.ListEstados).Methods("GET")
	router.HandleFunc("/api/cidades", handlers.ListCidades).Methods("GET")
	router.HandleFunc("/api/generos", handlers.ListGeneros).Methods("GET")
	router.HandleFunc("/api/tipos", handlers.ListTipos).Methods("GET")
	router.HandleFunc("/api/classificacoes-etarias", handlers.ListClassificacoesEtarias).Methods("GET")
	router.HandleFunc("/api/classificacoes-internas", handlers.ListClassificacoesInternas).Methods("GET")
	// Adicionar as rotas dos cadastros,
	// Exemplo:
	// router.HandleFunc("/atores", AtorHandler).Methods("GET", "POST")

	return router
}
