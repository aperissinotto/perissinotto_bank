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
		INSERT INTO contas (id, cliente_id, descricao)
		VALUES ($1, $2, $3)
	`

	_, err := r.db.Exec(query, conta.ID, conta.ClienteId, conta.Descricao)
	return err
}

func (r *ContaRepositoryPostgres) BuscarContasPorClienteID(cliente_id string) ([]entity.Conta, error) {
	query := `
		SELECT id, cliente_id, descricao
		FROM contas
		WHERE cliente_id = $1
	`

	rows, err := r.db.Query(query, cliente_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contas []entity.Conta

	for rows.Next() {
		var c entity.Conta
		if err := rows.Scan(
			&c.ID,
			&c.ClienteId,
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
		SELECT id, cliente_id, descricao
		FROM contas
		WHERE id = $1`,
		id,
	).Scan(&c.ID, &c.ClienteId, &c.Descricao)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
