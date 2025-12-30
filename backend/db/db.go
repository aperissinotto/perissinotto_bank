package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func ConectarPostgres() {
	var err error

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	//dsn := "postgres://postgres:admin123@localhost:5432/postgres?sslmode=disable"

	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Erro ao abrir conexão:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Erro ao conectar no PostgreSQL:", err)
	}

	log.Println("PostgreSQL conectado")
}
