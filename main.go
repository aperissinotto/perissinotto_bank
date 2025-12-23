package main

import (
	"log"
	"net/http"
)

func main() {
	ConectarPostgres()
	defer DB.Close()
	// Servir arquivos estáticos da pasta public
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// Rotas de API
	http.HandleFunc("/api/login", HandleLogin)
	http.HandleFunc("/api/cadastrar", HandleCadastro)

	// Iniciar servidor
	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
