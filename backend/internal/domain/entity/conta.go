package entity

import "github.com/google/uuid"

type Conta struct {
	ID        uuid.UUID
	CPF       string
	Descricao string
}
