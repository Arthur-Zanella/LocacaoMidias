package main

import (
	"locacaomidias/internal/database"
	"locacaomidias/internal/routes"
	"log"
	"net/http"
)

func main() {
	// Conecta ao banco
	database.Connect()
	// Configura rotas
	router := routes.InitRoutes()

	log.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
