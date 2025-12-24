package main

import (
	"log"
	"net/http"

	"github.com/aperissinotto/perissinotto_bank/db"
	"github.com/aperissinotto/perissinotto_bank/handlers"
)

func main() {
	db.ConectarPostgres()
	defer db.DB.Close()
	// Servir arquivos estáticos da pasta public
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// Rotas de API
	http.HandleFunc("/api/login", handlers.HandleLogin)
	http.HandleFunc("/api/cadastrar", handlers.HandleCadastro)

	// Iniciar servidor
	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
