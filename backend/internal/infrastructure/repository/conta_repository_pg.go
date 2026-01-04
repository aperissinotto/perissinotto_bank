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

func (r *ContaRepositoryPostgres) CriarConta(conta *entity.Conta) error {
	query := `
		INSERT INTO contas (id, cpf, descricao)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(query, conta.ID, conta.CPF, conta.Descricao)
	return err
}

func (r *ContaRepositoryPostgres) BuscarContasPorCpf(cpf string) ([]entity.Conta, error) {
	query := `
		SELECT id, cpf, descricao
		FROM contas
		WHERE cpf = $1
	`

	rows, err := r.db.Query(query, cpf)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contas []entity.Conta

	for rows.Next() {
		var c entity.Conta
		if err := rows.Scan(
			&c.ID,
			&c.CPF,
			&c.Descricao,
		); err != nil {
			return nil, err
		}

		contas = append(contas, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return contas, nil
}

func (r *ContaRepositoryPostgres) BuscarContaPorId(id string) (*entity.Conta, error) {
	var c entity.Conta

	err := r.db.QueryRow(`
		SELECT id, cpf, descricao
		FROM contas
		WHERE id = $1`,
		id,
	).Scan(&c.ID, &c.CPF, &c.Descricao)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
