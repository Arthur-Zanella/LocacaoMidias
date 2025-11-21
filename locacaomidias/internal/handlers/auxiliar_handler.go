package handlers

import (
	"encoding/json"
	"locacaomidias/internal/database"
	"locacaomidias/internal/models"
	"log"
	"net/http"
)

// ===== ESTADOS =====
func ListEstados(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, nome, sigla FROM estado ORDER BY nome")
	if err != nil {
		log.Println("Erro ao buscar estados:", err)
		http.Error(w, "Erro ao buscar estados", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var estados []models.Estado
	for rows.Next() {
		var e models.Estado
		if err := rows.Scan(&e.ID, &e.Nome, &e.Sigla); err != nil {
			log.Println("Erro ao escanear estado:", err)
			continue
		}
		estados = append(estados, e)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(estados)
}

// ===== CIDADES =====
func ListCidades(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, nome, estado_id FROM cidade ORDER BY nome")
	if err != nil {
		log.Println("Erro ao buscar cidades:", err)
		http.Error(w, "Erro ao buscar cidades", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var cidades []models.Cidade
	for rows.Next() {
		var c models.Cidade
		if err := rows.Scan(&c.ID, &c.Nome, &c.EstadoID); err != nil {
			log.Println("Erro ao escanear cidade:", err)
			continue
		}
		cidades = append(cidades, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cidades)
}

// ===== GÊNEROS =====
func ListGeneros(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, descricao FROM genero ORDER BY descricao")
	if err != nil {
		log.Println("Erro ao buscar gêneros:", err)
		http.Error(w, "Erro ao buscar gêneros", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var generos []models.Genero
	for rows.Next() {
		var g models.Genero
		if err := rows.Scan(&g.ID, &g.Descricao); err != nil {
			log.Println("Erro ao escanear gênero:", err)
			continue
		}
		generos = append(generos, g)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(generos)
}

// ===== TIPOS =====
func ListTipos(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, descricao FROM tipo ORDER BY descricao")
	if err != nil {
		log.Println("Erro ao buscar tipos:", err)
		http.Error(w, "Erro ao buscar tipos", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tipos []models.Tipo
	for rows.Next() {
		var t models.Tipo
		if err := rows.Scan(&t.ID, &t.Descricao); err != nil {
			log.Println("Erro ao escanear tipo:", err)
			continue
		}
		tipos = append(tipos, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tipos)
}

// ===== CLASSIFICAÇÕES ETÁRIAS =====
func ListClassificacoesEtarias(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, descricao FROM classificacao_etaria ORDER BY id")
	if err != nil {
		log.Println("Erro ao buscar classificações etárias:", err)
		http.Error(w, "Erro ao buscar classificações etárias", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var classificacoes []models.ClassificacaoEtaria
	for rows.Next() {
		var c models.ClassificacaoEtaria
		if err := rows.Scan(&c.ID, &c.Descricao); err != nil {
			log.Println("Erro ao escanear classificação etária:", err)
			continue
		}
		classificacoes = append(classificacoes, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classificacoes)
}

// ===== CLASSIFICAÇÕES INTERNAS =====
func ListClassificacoesInternas(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, descricao, valor_aluguel FROM classificacao_interna ORDER BY descricao")
	if err != nil {
		log.Println("Erro ao buscar classificações internas:", err)
		http.Error(w, "Erro ao buscar classificações internas", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var classificacoes []models.ClassificacaoInterna
	for rows.Next() {
		var c models.ClassificacaoInterna
		if err := rows.Scan(&c.ID, &c.Descricao, &c.ValorAluguel); err != nil {
			log.Println("Erro ao escanear classificação interna:", err)
			continue
		}
		classificacoes = append(classificacoes, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(classificacoes)
}
