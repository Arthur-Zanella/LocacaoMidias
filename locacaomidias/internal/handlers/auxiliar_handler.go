package handlers

import (
	"encoding/json"
	"locacaomidias/internal/database"
	"locacaomidias/internal/models"
	"net/http"
)

// estados
func ListEstados(w http.ResponseWriter, r *http.Request) {
	var estados []models.Estado

	result := database.DB.Order("nome").Find(&estados)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(estados)
}

// cidades
func ListCidades(w http.ResponseWriter, r *http.Request) {
	var cidades []models.Cidade

	// Preload carrega o relacionamento com Estado
	result := database.DB.Preload("Estado").Order("nome").Find(&cidades)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cidades)
}

// generos
func ListGeneros(w http.ResponseWriter, r *http.Request) {
	var generos []models.Genero

	result := database.DB.Order("descricao").Find(&generos)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(generos)
}

// tipos
func ListTipos(w http.ResponseWriter, r *http.Request) {
	var tipos []models.Tipo

	result := database.DB.Order("descricao").Find(&tipos)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tipos)
}

// classificacoes etarias
func ListClassificacoesEtarias(w http.ResponseWriter, r *http.Request) {
	var classificacoes []models.ClassificacaoEtaria

	result := database.DB.Order("id").Find(&classificacoes)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classificacoes)
}

// classificacoes internas
func ListClassificacoesInternas(w http.ResponseWriter, r *http.Request) {
	var classificacoes []models.ClassificacaoInterna

	result := database.DB.Order("descricao").Find(&classificacoes)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classificacoes)
}
