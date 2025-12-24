package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aperissinotto/perissinotto_bank/auth"
	"github.com/aperissinotto/perissinotto_bank/models"
	"github.com/aperissinotto/perissinotto_bank/repository"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(models.LoginResponse{
			Success: false,
			Message: "Apenas POST é permitido",
		})
		return
	}

	var req models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.LoginResponse{
			Success: false,
			Message: "JSON inválido",
		})
		return
	}

	if req.Agencia == "" || req.Conta == "" || req.Senha == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.LoginResponse{
			Success: false,
			Message: "Agência, conta e senha são obrigatórios",
		})
		return
	}

	log.Printf("Login tentado - Agência: %s, Conta: %s", req.Agencia, req.Conta)
	// aqui vou acessar o postgresql para validar a senha
	idConta, senhaHash, err := repository.BuscarContaPorAgenciaConta(req.Agencia, req.Conta)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.LoginResponse{
			Success: false,
			Message: "Agência, conta ou senha inválidos",
		})
		return
	}

	if !auth.ValidarSenha(senhaHash, req.Senha) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.LoginResponse{
			Success: false,
			Message: "Agência, conta ou senha inválidos",
		})
		return
	}

	// cria sessão usando UUID da conta
	auth.CriarSessao(w, idConta)
	log.Println(idConta)
	// aqui vou acessar o postgresql para validar a senha

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.LoginResponse{
		Success: true,
		Message: "Login realizado com sucesso",
		Token:   "token_aqui",
	})
}

func HandleCadastro(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(models.LoginResponse{
			Success: false,
			Message: "Apenas POST é permitido",
		})
		return
	}

	var req models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.LoginResponse{
			Success: false,
			Message: "JSON inválido",
		})
		return
	}

	if req.Agencia == "" || req.Conta == "" || req.Senha == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.LoginResponse{
			Success: false,
			Message: "Agência, conta e senha são obrigatórios",
		})
		return
	}

	log.Printf("Login tentado - Agência: %s, Conta: %s", req.Agencia, req.Conta)
	// aqui vou acessar o postgresql para criar a conta
	erro := repository.CriarConta(req.Agencia, req.Conta, req.Senha)
	if erro != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.LoginResponse{
			Success: false,
			Message: "Agência, conta ou senha inválidos",
		})
		return
	}
	// aqui vou acessar o postgresql para criar a conta

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.LoginResponse{
		Success: true,
		Message: "Conta criada com sucesso",
		Token:   "token_aqui",
	})
}
