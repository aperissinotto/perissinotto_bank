package dto

type CriarClienteResponse struct {
	ID           string `json:"id"`
	NomeCompleto string `json:"nomeCompleto"`
	Email        string `json:"email"`
	CPF          string `json:"cpf"`
}
