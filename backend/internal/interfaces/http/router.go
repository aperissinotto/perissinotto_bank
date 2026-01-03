package http

import (
	"net/http"

	"github.com/aperissinotto/perissinotto_bank/internal/interfaces/http/handler"
)

func NewRouter(
	clienteHandler *handler.ClienteHandler,
) http.Handler {

	mux := http.NewServeMux()

	// Rotas de clientes
	mux.HandleFunc("/api/clientes", clienteHandler.CriarCliente)

	return mux
}
