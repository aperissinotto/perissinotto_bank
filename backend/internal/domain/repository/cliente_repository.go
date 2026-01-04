package repository

import "github.com/aperissinotto/perissinotto_bank/internal/domain/entity"

type ClienteRepository interface {
	CriarCliente(cliente *entity.Cliente) error
	BuscarClientePorCpf(cpf string) (*entity.Cliente, error)
}
