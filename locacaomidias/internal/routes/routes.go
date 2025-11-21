package routes

import (
	"github.com/gorilla/mux"
)

// InitRoutes inicializa o router da aplicação
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	// Adicionar as rotas dos cadastros,
	// Exemplo:
	// router.HandleFunc("/atores", AtorHandler).Methods("GET", "POST")

	return router
}
