package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"locacaomidias/internal/database"
	"locacaomidias/internal/models"

	"github.com/gorilla/mux"
)

// ListAtores retorna todos os atores (ordenados por nome)
func ListAtores(w http.ResponseWriter, r *http.Request) {
	var atores []models.Ator
	result := database.DB.Order("id").Find(&atores)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(atores)
}

// GetAtor busca ator por id
func GetAtor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, _ := strconv.Atoi(idStr)

	var ator models.Ator
	result := database.DB.First(&ator, id)
	if result.Error != nil {
		http.Error(w, "ator not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ator)
}

// actorInput usado para permitir data no formato "YYYY-MM-DD"
type actorInput struct {
	Nome        string `json:"nome"`
	Sobrenome   string `json:"sobrenome"`
	DataEstreia string `json:"data_estreia"` // aceita "2006-01-02"
}

// CreateAtor cria um novo ator
func CreateAtor(w http.ResponseWriter, r *http.Request) {
	var input actorInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	var data time.Time
	if input.DataEstreia != "" {
		t, err := time.Parse("2006-01-02", input.DataEstreia)
		if err != nil {
			http.Error(w, "invalid date format, use YYYY-MM-DD", http.StatusBadRequest)
			return
		}
		data = t
	} else {
		data = time.Time{} // zero value
	}

	ator := models.Ator{
		Nome:        input.Nome,
		Sobrenome:   input.Sobrenome,
		DataEstreia: data,
	}

	if err := database.DB.Create(&ator).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ator)
}

// UpdateAtor atualiza um ator existente
func UpdateAtor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var ator models.Ator
	if err := database.DB.First(&ator, id).Error; err != nil {
		http.Error(w, "ator not found", http.StatusNotFound)
		return
	}

	var input actorInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	ator.Nome = input.Nome
	ator.Sobrenome = input.Sobrenome
	if input.DataEstreia != "" {
		t, err := time.Parse("2006-01-02", input.DataEstreia)
		if err != nil {
			http.Error(w, "invalid date format, use YYYY-MM-DD", http.StatusBadRequest)
			return
		}
		ator.DataEstreia = t
	}

	if err := database.DB.Save(&ator).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(ator)
}

// DeleteAtor deleta um ator por id
func DeleteAtor(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := database.DB.Delete(&models.Ator{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
