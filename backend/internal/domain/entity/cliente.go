package entity

import "github.com/google/uuid"

type Cliente struct {
	ID             uuid.UUID
	NomeCompleto   string
	Email          string
	DataNascimento string
	CPF            string
	RG             string
	CEP            string
	Endereco       string
	Bairro         string
	Cidade         string
	Estado         string
	RendaMensal    float64
}
