package models

// Request de login vindo do frontend
type LoginRequest struct {
	Agencia string `json:"agencia"`
	Conta   string `json:"conta"`
	Senha   string `json:"senha"`
}

// Resposta do login
type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Token   string `json:"token,omitempty"`
}

// Estrutura de uma conta no banco
type Conta struct {
	ID      string `json:"id"`
	Agencia string `json:"agencia"`
	Conta   string `json:"conta"`
	Senha   string `json:"-"`
}
