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

func (r *ClienteRepositoryPostgres) CriarCliente(c *entity.Cliente) error {
	query := `
		INSERT INTO clientes 
		(nome_completo, email, data_nascimento, cpf, rg, cep, endereco, bairro, cidade, estado, renda_mensal, senha_hash)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
		RETURNING id
	`
	return r.db.QueryRow(
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
		c.SenhaHash,
	).Scan(&c.ID)
}

func (r *ClienteRepositoryPostgres) BuscarClientePorCpf(cpf string) (*entity.Cliente, error) {
	var c entity.Cliente

	err := r.db.QueryRow(`
		SELECT id, nome_completo, email, data_nascimento, cpf, rg,
		       cep, endereco, bairro, cidade, estado, renda_mensal, senha_hash
		FROM clientes
		WHERE cpf = $1
	`, cpf).Scan(
		&c.ID, &c.NomeCompleto, &c.Email, &c.DataNascimento,
		&c.CPF, &c.RG, &c.CEP, &c.Endereco,
		&c.Bairro, &c.Cidade, &c.Estado, &c.RendaMensal, &c.SenhaHash,
	)

	return &c, err
}
