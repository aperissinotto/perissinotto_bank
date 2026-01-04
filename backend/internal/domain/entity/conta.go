package entity

import "github.com/google/uuid"

type Conta struct {
	ID        uuid.UUID
	ClienteId uuid.UUID
	Descricao string
}
