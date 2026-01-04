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
	res := json.NewDecoder(r.Body).Decode(&req)
	if res != nil {
		WriteError(
			w,
			http.StatusUnauthorized,
			"JSON inválido!",
			res.Error(),
		)
		return
	}

	token, err := h.auth.Login(req.Cpf, req.Senha)

	if err != nil {
		WriteError(
			w,
			http.StatusUnauthorized,
			"CPF ou Senha inválidos!",
			err.Error(),
		)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   7200,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"login realizado com sucesso"}`))
}
