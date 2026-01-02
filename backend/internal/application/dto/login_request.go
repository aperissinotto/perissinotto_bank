package dto

type LoginRequest struct {
	Agencia string `json:"agencia"`
	Conta   string `json:"conta"`
	Senha   string `json:"senha"`
}
