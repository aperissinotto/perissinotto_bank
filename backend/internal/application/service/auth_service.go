package service

import (
	"errors"

	"github.com/aperissinotto/perissinotto_bank/internal/domain/entity"
	"github.com/aperissinotto/perissinotto_bank/internal/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.ContaRepository
}

func NewAuthService(r repository.ContaRepository) *AuthService {
	return &AuthService{repo: r}
}

func (s *AuthService) Login(agencia, conta, senha string) (*entity.Conta, error) {
	c, err := s.repo.BuscarPorAgenciaConta(agencia, conta)
	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(c.Senha), []byte(senha)) != nil {
		return nil, errors.New("credenciais inválidas")
	}

	return c, nil
}
