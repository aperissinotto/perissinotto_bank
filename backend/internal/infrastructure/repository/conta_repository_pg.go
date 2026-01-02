package repository

import (
	"database/sql"

	"github.com/aperissinotto/perissinotto_bank/internal/domain/entity"
)

type ContaRepositoryPostgres struct {
	db *sql.DB
}

func NewContaRepository(db *sql.DB) *ContaRepositoryPostgres {
	return &ContaRepositoryPostgres{db: db}
}

func (r *ContaRepositoryPostgres) BuscarPorAgenciaConta(agencia, conta string) (*entity.Conta, error) {
	var c entity.Conta

	err := r.db.QueryRow(`
		SELECT id, agencia, conta, senha
		FROM contas
		WHERE agencia = $1 AND conta = $2`,
		agencia, conta,
	).Scan(&c.ID, &c.Agencia, &c.Conta, &c.Senha)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *ContaRepositoryPostgres) Criar(conta *entity.Conta) error {
	query := `
		INSERT INTO contas (agencia, conta, senha)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(query, conta.Agencia, conta.Conta, conta.Senha)
	return err
}
