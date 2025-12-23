package main

import (
	"database/sql"
	"errors"
)

func BuscarContaPorAgenciaConta(agencia, conta string) (string, string, error) {
	var id string
	var senhaHash string

	err := DB.QueryRow(`
		SELECT id, senha
		FROM contas
		WHERE agencia = $1::numeric
		  AND conta   = $2::numeric
	`, agencia, conta).Scan(&id, &senhaHash)

	if err == sql.ErrNoRows {
		return "", "", errors.New("conta não encontrada")
	}

	if err != nil {
		return "", "", err
	}

	return id, senhaHash, nil
}
