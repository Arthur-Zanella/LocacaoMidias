package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	// variaveis do .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Não foi possível carregar o .env, usando variáveis do sistema")
	}

	// dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// abre conexao
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Log de queries
	})
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %s", err)
	}

	// testa conexao
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Erro ao obter conexão SQL: %s", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Erro ao fazer ping no banco: %s", err)
	}

	log.Println("Banco de dados conectado")
}
