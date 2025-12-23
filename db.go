package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func ConectarPostgres() {
	var err error

	dsn := "postgres://postgres:admin123@localhost:5432/postgres?sslmode=disable"

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
