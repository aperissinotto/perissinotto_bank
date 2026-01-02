package repository

import "github.com/aperissinotto/perissinotto_bank/internal/domain/entity"

type ClienteRepository interface {
	Criar(cliente *entity.Cliente) error
	BuscarPorID(id string) (*entity.Cliente, error)
}
