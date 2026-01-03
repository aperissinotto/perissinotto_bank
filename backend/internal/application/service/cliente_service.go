package service

import (
	"errors"

	"github.com/aperissinotto/perissinotto_bank/internal/application/dto"
	"github.com/aperissinotto/perissinotto_bank/internal/domain/entity"
	"github.com/aperissinotto/perissinotto_bank/internal/domain/repository"
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
	}

	// aqui entram regras de negócio:
	// - validar CPF
	// - validar idade
	// - hash da senha
	// - gerar ID

	if err := s.repo.Criar(cliente); err != nil {
		return nil, err
	}

	return cliente, nil
}

func (s *ClienteService) BuscarCliente(id string) (*entity.Cliente, error) {
	return s.repo.BuscarPorID(id)
}
