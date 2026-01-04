package service

import (
	"errors"
	"log"

	"github.com/aperissinotto/perissinotto_bank/internal/domain/entity"
	"github.com/aperissinotto/perissinotto_bank/internal/domain/repository"
	"github.com/aperissinotto/perissinotto_bank/internal/domain/security"
)

type LoginService struct {
	repoCliente repository.ClienteRepository
}

func NewLoginService(r repository.ClienteRepository) *LoginService {
	return &LoginService{repoCliente: r}
}

func (s *LoginService) Login(cpf, senhaAberta string) (*entity.Cliente, error) {
	c, err := s.repoCliente.BuscarClientePorCpf(cpf)
	if err != nil {
		log.Println(err)
		return nil, errors.New("001-credenciais inválidas")
	}

	res := security.CompararSenha(senhaAberta, c.SenhaHash)
	if !res {
		return nil, errors.New("002-credenciais inválidas")
	}

	return c, nil
}
