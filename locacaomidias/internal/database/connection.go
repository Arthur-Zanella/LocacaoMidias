package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {
	// Carrega variáveis do .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Não foi possível carregar o .env, usando variáveis do sistema")
	}

	// Monta o DSN (string de conexão)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Abre a conexão
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erro ao abrir conexão: %s", err)
	}

	// Testa conexão
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %s", err)
	}

	log.Println("Banco de dados conectado com sucesso!")
}
