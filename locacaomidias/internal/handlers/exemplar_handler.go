package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"locacaomidias/internal/database"
	"locacaomidias/internal/models"

	"github.com/gorilla/mux"
)

// ListExemplares lista todos os exemplares com preload da midia
func ListExemplares(w http.ResponseWriter, r *http.Request) {
	var exemplares []models.Exemplar
	result := database.DB.Preload("Midia.ClassificacaoInterna").Find(&exemplares)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exemplares)
}

// GetExemplar retorna um exemplar pelo codigo_interno
func GetExemplar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var exemplar models.Exemplar
	if err := database.DB.Preload("Midia").First(&exemplar, id).Error; err != nil {
		http.Error(w, "exemplar not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(exemplar)
}

// exemplarInput payload para criar/atualizar exemplar
type exemplarInput struct {
	Disponivel bool `json:"disponivel"`
	MidiaID    int  `json:"midia_id"`
}

// CreateExemplar cria um exemplar
func CreateExemplar(w http.ResponseWriter, r *http.Request) {
	var input exemplarInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	exemplar := models.Exemplar{
		Disponivel: input.Disponivel,
		MidiaID:    input.MidiaID,
	}

	if err := database.DB.Create(&exemplar).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database.DB.Preload("Midia").First(&exemplar, exemplar.CodigoInterno)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(exemplar)
}

// UpdateExemplar atualiza um exemplar existente
func UpdateExemplar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var exemplar models.Exemplar
	if err := database.DB.First(&exemplar, id).Error; err != nil {
		http.Error(w, "exemplar not found", http.StatusNotFound)
		return
	}

	var input exemplarInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	exemplar.Disponivel = input.Disponivel
	exemplar.MidiaID = input.MidiaID

	if err := database.DB.Save(&exemplar).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database.DB.Preload("Midia").First(&exemplar, exemplar.CodigoInterno)
	json.NewEncoder(w).Encode(exemplar)
}

// DeleteExemplar deleta exemplar por codigo_interno
func DeleteExemplar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := database.DB.Delete(&models.Exemplar{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
