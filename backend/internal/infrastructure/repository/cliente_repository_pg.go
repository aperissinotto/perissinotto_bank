package repository

import (
	"database/sql"

	"github.com/aperissinotto/perissinotto_bank/internal/domain/entity"
)

type ClienteRepositoryPostgres struct {
	db *sql.DB
}

func NewClienteRepository(db *sql.DB) *ClienteRepositoryPostgres {
	return &ClienteRepositoryPostgres{db: db}
}

func (r *ClienteRepositoryPostgres) Criar(c *entity.Cliente) error {
	query := `
		INSERT INTO clientes 
		(nome_completo, email, data_nascimento, cpf, rg, cep, endereco, bairro, cidade, estado, renda_mensal)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
	`
	_, err := r.db.Exec(
		query,
		c.NomeCompleto,
		c.Email,
		c.DataNascimento,
		c.CPF,
		c.RG,
		c.CEP,
		c.Endereco,
		c.Bairro,
		c.Cidade,
		c.Estado,
		c.RendaMensal,
	)
	return err
}

func (r *ClienteRepositoryPostgres) BuscarPorID(id string) (*entity.Cliente, error) {
	var c entity.Cliente

	err := r.db.QueryRow(`
		SELECT id, nome_completo, email, data_nascimento, cpf, rg,
		       cep, endereco, bairro, cidade, estado, renda_mensal
		FROM clientes
		WHERE id = $1
	`, id).Scan(
		&c.ID, &c.NomeCompleto, &c.Email, &c.DataNascimento,
		&c.CPF, &c.RG, &c.CEP, &c.Endereco,
		&c.Bairro, &c.Cidade, &c.Estado, &c.RendaMensal,
	)

	return &c, err
}
