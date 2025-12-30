package main

import (
	"log"
	"net/http"

	"github.com/aperissinotto/perissinotto_bank/handlers"
)

func main() {
	//	db.ConectarPostgres()
	//	defer db.DB.Close()

	// Rotas de API
	http.HandleFunc("/api/login", handlers.HandleLogin)
	http.HandleFunc("/api/cadastrar", handlers.HandleCadastro)

	// Iniciar servidor
	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
