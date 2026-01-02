package service

import (
	"github.com/aperissinotto/perissinotto_bank/internal/domain/entity"
	"github.com/aperissinotto/perissinotto_bank/internal/domain/repository"
)

type ClienteService struct {
	repo repository.ClienteRepository
}

func NewClienteService(repo repository.ClienteRepository) *ClienteService {
	return &ClienteService{repo: repo}
}

func (s *ClienteService) CriarCliente(c *entity.Cliente) error {
	return s.repo.Criar(c)
}

func (s *ClienteService) BuscarCliente(id string) (*entity.Cliente, error) {
	return s.repo.BuscarPorID(id)
}
