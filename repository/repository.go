package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/aperissinotto/perissinotto_bank/db"
	"golang.org/x/crypto/bcrypt"
)

func BuscarContaPorAgenciaConta(agencia, conta string) (string, string, error) {
	var id string
	var senhaHash string

	err := db.DB.QueryRow(`
		SELECT id, senha
		FROM contas
		WHERE agencia = $1::numeric
		  AND conta   = $2::numeric
	`, agencia, conta).Scan(&id, &senhaHash)

	if err == sql.ErrNoRows {
		log.Println("conta não encontrada!")
		return "", "", errors.New("conta não encontrada")
	}

	if err != nil {
		return "", "", err
	}

	return id, senhaHash, nil
}

func CriarConta(agencia, conta, senha string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)

	_, err := db.DB.Exec(`
		INSERT INTO contas (agencia, conta, senha)
		VALUES ($1, $2, $3)
	`, agencia, conta, hash)
	return err
}
