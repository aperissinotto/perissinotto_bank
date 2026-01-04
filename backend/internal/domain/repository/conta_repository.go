package repository

import "github.com/aperissinotto/perissinotto_bank/internal/domain/entity"

type ContaRepository interface {
	CriarConta(conta *entity.Conta) error
	BuscarContasPorCpf(cpf string) ([]entity.Conta, error)
	BuscarContaPorId(id string) (*entity.Conta, error)
}
