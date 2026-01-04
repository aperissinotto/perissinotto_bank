package dto

type LoginRequest struct {
	Cpf   string `json:"cpf"`
	Senha string `json:"senha"`
}
