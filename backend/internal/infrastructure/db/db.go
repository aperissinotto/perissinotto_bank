package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

// Connect inicializa a conexão com o PostgreSQL
func Connect() *sql.DB {
	if DB != nil {
		return DB
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Erro ao abrir conexão com o banco: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	log.Println("✅ Conectado ao PostgreSQL com sucesso")

	DB = db
	return db
}
