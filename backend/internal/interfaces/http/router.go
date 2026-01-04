package http

import (
	"net/http"

	"github.com/aperissinotto/perissinotto_bank/internal/interfaces/http/handler"
)

func NewRouter(
	clienteHandler *handler.ClienteHandler,
	loginHandler *handler.LoginHandler,
) http.Handler {

	mux := http.NewServeMux()

	// Rotas de clientes
	mux.HandleFunc("/api/clientes", clienteHandler.CriarCliente)
	// Rota de login
	mux.HandleFunc("/api/login", loginHandler.Login)

	return mux
}
