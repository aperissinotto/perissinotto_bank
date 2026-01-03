package http

import (
	"encoding/json"
	"net/http"

	"github.com/aperissinotto/perissinotto_bank/internal/application/dto"
	"github.com/aperissinotto/perissinotto_bank/internal/application/service"
)

type ClienteHandler struct {
	service *service.ClienteService
}

func NewClienteHandler(s *service.ClienteService) *ClienteHandler {
	return &ClienteHandler{service: s}
}

func (h *ClienteHandler) CriarCliente(w http.ResponseWriter, r *http.Request) {
	var req dto.CriarClienteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	cliente, err := h.service.CriarCliente(req)

	if err != nil {
		http.Error(w, "Erro ao criar cliente", http.StatusInternalServerError)
		return
	}

	resp := dto.ClienteFromEntity(cliente)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
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
