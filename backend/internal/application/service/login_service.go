package service

import (
	"errors"

	"github.com/aperissinotto/perissinotto_bank/internal/domain/repository"
	"github.com/aperissinotto/perissinotto_bank/internal/infrastructure/auth"
)

type LoginService struct {
	repoCliente repository.ClienteRepository
}

func NewLoginService(r repository.ClienteRepository) *LoginService {
	return &LoginService{repoCliente: r}
}

func (s *LoginService) Login(cpf string, senhaAberta string) (string, error) {
	c, err := s.repoCliente.BuscarClientePorCpf(cpf)
	if err != nil {
		return "", errors.New("001-credenciais inválidas")
	}

	res := auth.CompararSenha(senhaAberta, c.SenhaHash)
	if !res {
		return "", errors.New("002-credenciais inválidas")
	}

	token, err := auth.GerarToken(c.ID.String(), c.CPF)
	if err != nil {
		return "", err
	}

	return token, nil
}
