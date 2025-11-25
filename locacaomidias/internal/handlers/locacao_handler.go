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

// parseDate helper
func mustParse(date string) time.Time {
	t, _ := time.Parse("2006-01-02", date)
	return t
}

type locacaoInput struct {
	DataInicio string `json:"data_inicio"`
	DataFim    string `json:"data_fim"`
	Cancelada  bool   `json:"cancelada"`
	ClienteID  int    `json:"cliente_id"`
	ExemplarID int    `json:"exemplar_codigo_interno"` // deve receber código do exemplar
}

// ---------------------------
// LISTAR TODAS
// ---------------------------
func ListLocacoes(w http.ResponseWriter, r *http.Request) {
	var locacoes []models.Locacao

	err := database.DB.
		Preload("Cliente").
		Preload("Cliente.Cidade").
		Preload("Cliente.Cidade.Estado").
		Preload("Itens").
		Preload("Itens.Exemplar").
		Preload("Itens.Exemplar.Midia").
		Preload("Itens.Exemplar.Midia.ClassificacaoInterna").
		Order("id ASC").
		Find(&locacoes).Error

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(locacoes)
}

// ---------------------------
// OBTER UMA
// ---------------------------
func GetLocacao(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var locacao models.Locacao

	err := database.DB.
		Preload("Cliente").
		Preload("Itens").
		Preload("Itens.Exemplar").
		Preload("Itens.Exemplar.Midia").
		Preload("Itens.Exemplar.Midia.ClassificacaoInterna").
		First(&locacao, id).Error

	if err != nil {
		http.Error(w, "locacao not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(locacao)
}

// ---------------------------
// CRIAR
// ---------------------------
func CreateLocacao(w http.ResponseWriter, r *http.Request) {
	var input locacaoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// obtém exemplar
	var exemplar models.Exemplar
	if err := database.DB.Preload("Midia.ClassificacaoInterna").
		First(&exemplar, input.ExemplarID).Error; err != nil {
		http.Error(w, "exemplar not found", http.StatusBadRequest)
		return
	}

	valor := exemplar.Midia.ClassificacaoInterna.ValorAluguel

	locacao := models.Locacao{
		DataInicio: mustParse(input.DataInicio),
		DataFim:    mustParse(input.DataFim),
		ClienteID:  input.ClienteID,
		Cancelada:  false,
	}

	if err := database.DB.Create(&locacao).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	item := models.ItemLocacao{
		LocacaoID:      locacao.ID,
		ExemplarCodigo: exemplar.CodigoInterno,
		Valor:          valor,
	}

	if err := database.DB.Create(&item).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// recarrega locação completa
	database.DB.
		Preload("Itens").
		Preload("Itens.Exemplar").
		Preload("Itens.Exemplar.Midia").
		Preload("Itens.Exemplar.Midia.ClassificacaoInterna").
		Preload("Cliente").
		First(&locacao, locacao.ID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(locacao)
}

// ---------------------------
// ATUALIZAR
// ---------------------------
func UpdateLocacao(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var locacao models.Locacao
	if err := database.DB.First(&locacao, id).Error; err != nil {
		http.Error(w, "locacao not found", http.StatusNotFound)
		return
	}

	var input locacaoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// Atualiza campos editáveis
	locacao.DataInicio = mustParse(input.DataInicio)
	locacao.DataFim = mustParse(input.DataFim)
	locacao.Cancelada = input.Cancelada
	// não alteramos ClienteID porque não muda na edição

	// Atualiza item
	var item models.ItemLocacao
	database.DB.Where("locacao_id = ?", locacao.ID).First(&item)

	if input.ExemplarID != 0 && item.ExemplarCodigo != input.ExemplarID {
		var exemplar models.Exemplar
		if err := database.DB.Preload("Midia.ClassificacaoInterna").
			First(&exemplar, input.ExemplarID).Error; err == nil {
			item.ExemplarCodigo = exemplar.CodigoInterno
			item.Valor = exemplar.Midia.ClassificacaoInterna.ValorAluguel
		}
	}

	database.DB.Save(&locacao)
	database.DB.Save(&item)

	// recarrega locação completa
	database.DB.
		Preload("Itens").
		Preload("Itens.Exemplar").
		Preload("Itens.Exemplar.Midia").
		Preload("Itens.Exemplar.Midia.ClassificacaoInterna").
		Preload("Cliente").
		First(&locacao, id)

	json.NewEncoder(w).Encode(locacao)
}

// ---------------------------
// DELETAR
// ---------------------------
func DeleteLocacao(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	database.DB.Where("locacao_id = ?", id).Delete(&models.ItemLocacao{})
	database.DB.Delete(&models.Locacao{}, id)

	w.WriteHeader(http.StatusNoContent)
}
