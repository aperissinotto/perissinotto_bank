package handler

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

	if r.Header.Get("Content-Type") != "application/json" {
		writeError(
			w,
			http.StatusUnsupportedMediaType,
			"Content-Type deve ser application/json",
			"",
		)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(
			w,
			http.StatusBadRequest,
			"JSON inválido",
			err.Error(),
		)
		return
	}

	cliente, err := h.service.CriarCliente(req)

	if err != nil {
		writeError(
			w,
			http.StatusInternalServerError,
			"Erro ao criar cliente",
			err.Error(),
		)
		return
	}

	clienteResp := dto.ClienteFromEntity(cliente)

	response := dto.ApiResponse[dto.CriarClienteResponse]{
		Message: "Cadastro realizado com sucesso",
		Data:    clienteResp,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
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
