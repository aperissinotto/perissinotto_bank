package main

import (
	"log"
	"net/http"

	"github.com/aperissinotto/perissinotto_bank/internal/application/service"
	"github.com/aperissinotto/perissinotto_bank/internal/infrastructure/db"
	"github.com/aperissinotto/perissinotto_bank/internal/infrastructure/repository"
	httpHandler "github.com/aperissinotto/perissinotto_bank/internal/interfaces/http"
)

func main() {
	dbConn := db.Connect()

	contaRepo := repository.NewContaRepository(dbConn)
	clienteRepo := repository.NewClienteRepository(dbConn)

	contaService := service.NewAuthService(contaRepo)
	clienteService := service.NewClienteService(clienteRepo)

	contaHandler := httpHandler.NewHandler(contaService)
	clienteHandler := httpHandler.NewClienteHandler(clienteService)

	http.HandleFunc("/login", contaHandler.Login)
	http.HandleFunc("/clientes", clienteHandler.CriarCliente)
	http.HandleFunc("/clientes/buscar", clienteHandler.BuscarCliente)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
