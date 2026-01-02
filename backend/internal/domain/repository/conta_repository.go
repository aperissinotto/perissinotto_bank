package repository

import "github.com/aperissinotto/perissinotto_bank/internal/domain/entity"

type ContaRepository interface {
	BuscarPorAgenciaConta(agencia, conta string) (*entity.Conta, error)
	Criar(conta *entity.Conta) error
}
