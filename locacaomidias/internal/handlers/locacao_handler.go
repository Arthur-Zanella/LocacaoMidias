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
	ExemplarID int    `json:"exemplar_codigo_interno"`
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

	w.Header().Set("Content-Type", "application/json")
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locacao)
}

// ---------------------------
// CRIAR LOCAÇÃO
// ---------------------------
func CreateLocacao(w http.ResponseWriter, r *http.Request) {
	var input locacaoInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// Parse datas
	dataInicio := mustParse(input.DataInicio)
	dataFim := mustParse(input.DataFim)

	// Validar datas
	if dataFim.Before(dataInicio) {
		http.Error(w, "data_fim deve ser posterior a data_inicio", http.StatusBadRequest)
		return
	}

	// 1. VERIFICAR SE EXEMPLAR EXISTE
	var exemplar models.Exemplar
	if err := database.DB.Preload("Midia.ClassificacaoInterna").
		First(&exemplar, input.ExemplarID).Error; err != nil {
		http.Error(w, "exemplar not found", http.StatusNotFound)
		return
	}

	// 2. VERIFICAR SE EXEMPLAR ESTÁ DISPONÍVEL (só se locação NÃO for cancelada)
	if !input.Cancelada && !exemplar.Disponivel {
		http.Error(w, "exemplar não está disponível para locação", http.StatusConflict)
		return
	}

	// 3. CRIAR LOCAÇÃO
	valor := exemplar.Midia.ClassificacaoInterna.ValorAluguel

	locacao := models.Locacao{
		DataInicio: dataInicio,
		DataFim:    dataFim,
		ClienteID:  input.ClienteID,
		Cancelada:  input.Cancelada,
	}

	if err := database.DB.Create(&locacao).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 4. CRIAR ITEM DE LOCAÇÃO
	item := models.ItemLocacao{
		LocacaoID:      locacao.ID,
		ExemplarCodigo: exemplar.CodigoInterno,
		Valor:          valor,
	}

	if err := database.DB.Create(&item).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 5. MARCAR EXEMPLAR COMO INDISPONÍVEL (só se locação NÃO for cancelada)
	if !input.Cancelada {
		exemplar.Disponivel = false
		if err := database.DB.Save(&exemplar).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// 6. RECARREGAR LOCAÇÃO COMPLETA
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
// ATUALIZAR LOCAÇÃO
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

	// Parse datas
	dataInicio := mustParse(input.DataInicio)
	dataFim := mustParse(input.DataFim)

	// Validar datas
	if dataFim.Before(dataInicio) {
		http.Error(w, "data_fim deve ser posterior a data_inicio", http.StatusBadRequest)
		return
	}

	// Buscar item atual
	var item models.ItemLocacao
	database.DB.Where("locacao_id = ?", locacao.ID).First(&item)

	// VERIFICAR SE MUDOU O STATUS DE CANCELAMENTO
	statusCanceladoMudou := locacao.Cancelada != input.Cancelada

	// Se mudou o exemplar
	if input.ExemplarID != 0 && item.ExemplarCodigo != input.ExemplarID {
		// 1. VERIFICAR SE NOVO EXEMPLAR EXISTE
		var novoExemplar models.Exemplar
		if err := database.DB.Preload("Midia.ClassificacaoInterna").
			First(&novoExemplar, input.ExemplarID).Error; err != nil {
			http.Error(w, "novo exemplar not found", http.StatusNotFound)
			return
		}

		// 2. VERIFICAR SE NOVO EXEMPLAR ESTÁ DISPONÍVEL (só se locação NÃO for cancelada)
		if !input.Cancelada && !novoExemplar.Disponivel {
			http.Error(w, "novo exemplar não está disponível", http.StatusConflict)
			return
		}

		// 3. LIBERAR EXEMPLAR ANTIGO (só se locação antiga NÃO era cancelada)
		if !locacao.Cancelada {
			var exemplarAntigo models.Exemplar
			database.DB.First(&exemplarAntigo, item.ExemplarCodigo)
			exemplarAntigo.Disponivel = true
			database.DB.Save(&exemplarAntigo)
		}

		// 4. ATUALIZAR ITEM COM NOVO EXEMPLAR
		item.ExemplarCodigo = novoExemplar.CodigoInterno
		item.Valor = novoExemplar.Midia.ClassificacaoInterna.ValorAluguel

		// 5. BLOQUEAR NOVO EXEMPLAR (só se locação NÃO for cancelada)
		if !input.Cancelada {
			novoExemplar.Disponivel = false
			database.DB.Save(&novoExemplar)
		}
	} else if statusCanceladoMudou {
		// SE MUDOU APENAS O STATUS DE CANCELAMENTO (sem trocar exemplar)
		var exemplarAtual models.Exemplar
		database.DB.First(&exemplarAtual, item.ExemplarCodigo)

		if input.Cancelada {
			// Estava ativa, agora foi CANCELADA → liberar exemplar
			exemplarAtual.Disponivel = true
		} else {
			// Estava cancelada, agora foi ATIVADA → bloquear exemplar
			if !exemplarAtual.Disponivel {
				http.Error(w, "exemplar não está disponível para reativar locação", http.StatusConflict)
				return
			}
			exemplarAtual.Disponivel = false
		}
		database.DB.Save(&exemplarAtual)
	}

	// Atualizar campos da locação
	locacao.DataInicio = dataInicio
	locacao.DataFim = dataFim
	locacao.Cancelada = input.Cancelada

	// Salvar alterações
	database.DB.Save(&locacao)
	database.DB.Save(&item)

	// Recarregar locação completa
	database.DB.
		Preload("Itens").
		Preload("Itens.Exemplar").
		Preload("Itens.Exemplar.Midia").
		Preload("Itens.Exemplar.Midia.ClassificacaoInterna").
		Preload("Cliente").
		First(&locacao, id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(locacao)
}

// ---------------------------
// DELETAR LOCAÇÃO
// ---------------------------
func DeleteLocacao(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// Buscar locação
	var locacao models.Locacao
	if err := database.DB.First(&locacao, id).Error; err != nil {
		http.Error(w, "locacao not found", http.StatusNotFound)
		return
	}

	// Buscar itens da locação
	var itens []models.ItemLocacao
	database.DB.Where("locacao_id = ?", id).Find(&itens)

	// Liberar exemplares (só se a locação NÃO era cancelada)
	if !locacao.Cancelada {
		for _, item := range itens {
			var exemplar models.Exemplar
			if database.DB.First(&exemplar, item.ExemplarCodigo).Error == nil {
				exemplar.Disponivel = true
				database.DB.Save(&exemplar)
			}
		}
	}

	// Deletar itens e locação
	database.DB.Where("locacao_id = ?", id).Delete(&models.ItemLocacao{})
	database.DB.Delete(&models.Locacao{}, id)

	w.WriteHeader(http.StatusNoContent)
}
