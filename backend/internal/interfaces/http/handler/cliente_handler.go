package handler

import (
	"encoding/json"
	"net/http"
	"strings"

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

	if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
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

func (h *ClienteHandler) BuscarClientePorCpf(w http.ResponseWriter, r *http.Request) {
	cpf := r.URL.Query().Get("cpf")

	cliente, err := h.service.BuscarClientePorCpf(cpf)
	if err != nil {
		writeError(
			w,
			http.StatusNotFound,
			"Cliente não cadastrado!",
			err.Error(),
		)
		return
	}

	resp := dto.ClienteFromEntity(cliente)
	json.NewEncoder(w).Encode(resp)
}
