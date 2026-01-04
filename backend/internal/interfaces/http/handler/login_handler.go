package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aperissinotto/perissinotto_bank/internal/application/dto"
	"github.com/aperissinotto/perissinotto_bank/internal/application/service"
)

type LoginHandler struct {
	auth *service.LoginService
}

func NewLoginHandler(auth *service.LoginService) *LoginHandler {
	return &LoginHandler{auth: auth}
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	_, err := h.auth.Login(req.Cpf, req.Senha)

	if err != nil {
		writeError(
			w,
			http.StatusUnauthorized,
			"CPF ou Senha inválidos!",
			err.Error(),
		)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"login ok"}`))
}
