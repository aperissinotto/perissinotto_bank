package main

import (
	"log"
	"net/http"

	"github.com/aperissinotto/perissinotto_bank/internal/application/service"
	"github.com/aperissinotto/perissinotto_bank/internal/infrastructure/db"
	"github.com/aperissinotto/perissinotto_bank/internal/infrastructure/repository"
	httpRouter "github.com/aperissinotto/perissinotto_bank/internal/interfaces/http"
	"github.com/aperissinotto/perissinotto_bank/internal/interfaces/http/handler"
)

func main() {
	// Banco
	dbConn := db.Connect()

	// Repositories
	clienteRepo := repository.NewClienteRepository(dbConn)

	// Services
	clienteService := service.NewClienteService(clienteRepo)
	loginService := service.NewLoginService(clienteRepo)

	// Handlers
	clienteHandler := handler.NewClienteHandler(clienteService)
	loginHandler := handler.NewLoginHandler(loginService)

	// Router
	router := httpRouter.NewRouter(
		clienteHandler,
		loginHandler,
	)

	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
