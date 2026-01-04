package service

import (
	"errors"

	"github.com/aperissinotto/perissinotto_bank/internal/application/dto"
	"github.com/aperissinotto/perissinotto_bank/internal/domain/entity"
	"github.com/aperissinotto/perissinotto_bank/internal/domain/repository"
	"github.com/aperissinotto/perissinotto_bank/internal/domain/security"
	"github.com/aperissinotto/perissinotto_bank/internal/domain/validation"
)

type ClienteService struct {
	repo repository.ClienteRepository
}

func NewClienteService(repo repository.ClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

func (s *ClienteService) CriarCliente(req dto.CriarClienteRequest) (*entity.Cliente, error) {

	if !validation.ValidarCPF(req.CPF) {
		return nil, errors.New("CPF inválido")
	}

	if !validation.ValidarSenha(req.Senha) {
		return nil, errors.New("Senha inválida")
	}

	hash, res := security.HashSenha(req.Senha)
	if !res {
		return nil, errors.New("erro ao processar senha")
	}

	cliente := &entity.Cliente{
		NomeCompleto:   req.NomeCompleto,
		Email:          req.Email,
		DataNascimento: req.DataNascimento,
		CPF:            req.CPF,
		RG:             req.RG,
		CEP:            req.CEP,
		Endereco:       req.Endereco,
		Bairro:         req.Bairro,
		Cidade:         req.Cidade,
		Estado:         req.Estado,
		RendaMensal:    req.RendaMensal,
		SenhaHash:      hash,
	}

	if err := s.repo.CriarCliente(cliente); err != nil {
		return nil, err
	}

	return cliente, nil
}

func (s *ClienteService) BuscarClientePorCpf(cpf string) (*entity.Cliente, error) {
	return s.repo.BuscarClientePorCpf(cpf)
}
