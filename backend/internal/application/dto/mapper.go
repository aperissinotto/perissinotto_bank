package dto

import "github.com/aperissinotto/perissinotto_bank/internal/domain/entity"

func ClienteFromEntity(c *entity.Cliente) CriarClienteResponse {
	return CriarClienteResponse{
		ID:           c.ID.String(),
		NomeCompleto: c.NomeCompleto,
		Email:        c.Email,
		CPF:          c.CPF,
	}
}
