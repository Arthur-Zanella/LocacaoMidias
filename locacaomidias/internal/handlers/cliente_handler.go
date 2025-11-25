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

func parseDate(dateStr string) (time.Time, error) {
	layout := "2006-01-02"
	return time.Parse(layout, dateStr)
}

// ListClientes lista todos os clientes
func ListClientes(w http.ResponseWriter, r *http.Request) {
	var clientes []models.Cliente
	result := database.DB.
		Preload("Cidade").
		Preload("Cidade.Estado").
		Order("id ASC").
		Find(&clientes)

	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

// GetCliente retorna um cliente específico
func GetCliente(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var cliente models.Cliente
	result := database.DB.
		Preload("Cidade").
		Preload("Cidade.Estado").
		First(&cliente, id)

	if result.Error != nil {
		http.Error(w, "cliente not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cliente)
}

type clienteInput struct {
	Nome           string `json:"nome"`
	Sobrenome      string `json:"sobrenome"`
	DataNascimento string `json:"data_nascimento"`
	CPF            string `json:"cpf"`
	Email          string `json:"email"`
	Logradouro     string `json:"logradouro"`
	Numero         string `json:"numero"`
	Bairro         string `json:"bairro"`
	CEP            string `json:"cep"`
	CidadeID       int    `json:"cidade_id"`
}

// CreateCliente cria um novo cliente
func CreateCliente(w http.ResponseWriter, r *http.Request) {
	var input clienteInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// ---- FIX DO ERRO ----
	t, err := parseDate(input.DataNascimento)
	if err != nil {
		http.Error(w, "data_nascimento inválida (use YYYY-MM-DD)", http.StatusBadRequest)
		return
	}

	cliente := models.Cliente{
		Nome:           input.Nome,
		Sobrenome:      input.Sobrenome,
		DataNascimento: t,
		CPF:            input.CPF,
		Email:          input.Email,
		Logradouro:     input.Logradouro,
		Numero:         input.Numero,
		Bairro:         input.Bairro,
		CEP:            input.CEP,
		CidadeID:       input.CidadeID,
	}

	if err := database.DB.Create(&cliente).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database.DB.Preload("Cidade").Preload("Cidade.Estado").First(&cliente, cliente.ID)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cliente)
}

// UpdateCliente atualiza um cliente
func UpdateCliente(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var cliente models.Cliente
	if err := database.DB.First(&cliente, id).Error; err != nil {
		http.Error(w, "cliente not found", http.StatusNotFound)
		return
	}

	var input clienteInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid payload", http.StatusBadRequest)
		return
	}

	// ---- FIX DO ERRO ----
	t, err := parseDate(input.DataNascimento)
	if err != nil {
		http.Error(w, "data_nascimento inválida (use YYYY-MM-DD)", http.StatusBadRequest)
		return
	}

	cliente.Nome = input.Nome
	cliente.Sobrenome = input.Sobrenome
	cliente.DataNascimento = t
	cliente.CPF = input.CPF
	cliente.Email = input.Email
	cliente.Logradouro = input.Logradouro
	cliente.Numero = input.Numero
	cliente.Bairro = input.Bairro
	cliente.CEP = input.CEP
	cliente.CidadeID = input.CidadeID

	if err := database.DB.Save(&cliente).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	database.DB.Preload("Cidade").Preload("Cidade.Estado").First(&cliente, cliente.ID)

	json.NewEncoder(w).Encode(cliente)
}

// DeleteCliente remove um cliente
func DeleteCliente(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := database.DB.Delete(&models.Cliente{}, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
