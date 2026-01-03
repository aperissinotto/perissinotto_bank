package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aperissinotto/perissinotto_bank/internal/application/service"
	"github.com/aperissinotto/perissinotto_bank/internal/domain/entity"
)

type ClienteHandler struct {
	service *service.ClienteService
}

func NewClienteHandler(s *service.ClienteService) *ClienteHandler {
	return &ClienteHandler{service: s}
}

func (h *ClienteHandler) CriarCliente(w http.ResponseWriter, r *http.Request) {
	var c entity.Cliente
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err := h.service.CriarCliente(&c)
	log.Println(err)
	if err != nil {
		http.Error(w, "Erro ao criar cliente", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ClienteHandler) BuscarCliente(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	cliente, err := h.service.BuscarCliente(id)
	if err != nil {
		http.Error(w, "Cliente não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(cliente)
}
