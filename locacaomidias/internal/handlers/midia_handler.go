package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"locacaomidias/internal/database"
	"locacaomidias/internal/models"

	"github.com/gorilla/mux"
)

// ListMidias lista todas as mídias com preload dos relacionamentos
func ListMidias(w http.ResponseWriter, r *http.Request) {
	var midias []models.Midia
	result := database.DB.
		Preload("Genero").
		Preload("ClassificacaoEtaria").
		Preload("Tipo").
		Preload("ClassificacaoInterna").
		Preload("AtorPrincipal").
		Preload("AtorCoadjuvante").
		Order("id ASC").
		Find(&midias)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(midias)
}

// GetMidia retorna uma midia por id com relacionamentos
func GetMidia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var midia models.Midia
	result := database.DB.
		Preload("Genero").
		Preload("ClassificacaoEtaria").
		Preload("Tipo").
		Preload("ClassificacaoInterna").
		Preload("AtorPrincipal").
		Preload("AtorCoadjuvante").
		First(&midia, id)

	if result.Error != nil {
		http.Error(w, "midia not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(midia)
}

type midiaInput struct {
	Titulo                 string `json:"titulo"`
	AnoLancamento          string `json:"ano_lancamento"`
	CodigoBarras           string `json:"codigo_barras"`
	DuracaoEmMinutos       int    `json:"duracao_em_minutos"`
	AtorPrincipalID        int    `json:"ator_principal"`
	AtorCoadjuvanteID      int    `json:"ator_coadjuvante"`
	GeneroID               int    `json:"genero_id"`
	ClassificacaoEtariaID  int    `json:"classificacao_etaria_id"`
	TipoID                 int    `json:"tipo_id"`
	ClassificacaoInternaID int    `json:"classificacao_interna_id"`
}

// CreateMidia cria uma nova mídia
func CreateMidia(w http.ResponseWriter, r *http.Request) {
	var input midiaInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	midia := models.Midia{
		Titulo:                 input.Titulo,
		AnoLancamento:          input.AnoLancamento,
		CodigoBarras:           input.CodigoBarras,
		DuracaoEmMinutos:       input.DuracaoEmMinutos,
		AtorPrincipalID:        input.AtorPrincipalID,
		AtorCoadjuvanteID:      input.AtorCoadjuvanteID,
		GeneroID:               input.GeneroID,
		ClassificacaoEtariaID:  input.ClassificacaoEtariaID,
		TipoID:                 input.TipoID,
		ClassificacaoInternaID: input.ClassificacaoInternaID,
	}

	if err := database.DB.Create(&midia).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database.DB.Preload("Genero").
		Preload("ClassificacaoEtaria").
		Preload("Tipo").
		Preload("ClassificacaoInterna").
		Preload("AtorPrincipal").
		Preload("AtorCoadjuvante").
		First(&midia, midia.ID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(midia)
}

// UpdateMidia atualiza a midia existente
func UpdateMidia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var midia models.Midia
	if err := database.DB.First(&midia, id).Error; err != nil {
		http.Error(w, "midia not found", http.StatusNotFound)
		return
	}

	var input midiaInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	midia.Titulo = input.Titulo
	midia.AnoLancamento = input.AnoLancamento
	midia.CodigoBarras = input.CodigoBarras
	midia.DuracaoEmMinutos = input.DuracaoEmMinutos
	midia.AtorPrincipalID = input.AtorPrincipalID
	midia.AtorCoadjuvanteID = input.AtorCoadjuvanteID
	midia.GeneroID = input.GeneroID
	midia.ClassificacaoEtariaID = input.ClassificacaoEtariaID
	midia.TipoID = input.TipoID
	midia.ClassificacaoInternaID = input.ClassificacaoInternaID

	if err := database.DB.Save(&midia).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database.DB.Preload("Genero").
		Preload("ClassificacaoEtaria").
		Preload("Tipo").
		Preload("ClassificacaoInterna").
		Preload("AtorPrincipal").
		Preload("AtorCoadjuvante").
		First(&midia, midia.ID)

	json.NewEncoder(w).Encode(midia)
}

// DeleteMidia deleta uma midia por id
func DeleteMidia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if err := database.DB.Delete(&models.Midia{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
