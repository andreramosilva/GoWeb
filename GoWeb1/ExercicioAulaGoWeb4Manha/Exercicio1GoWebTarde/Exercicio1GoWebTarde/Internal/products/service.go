package products

// Service, deve conter a lógica da nossa aplicação.
// O arquivo service.go deve ser criado.
// A interface Service com todos os seus métodos deve ser gerada.
// A estrutura de serviço que contém o repositório deve ser gerada.
// Deve ser gerada uma função que retorne o Serviço.
// Todos os métodos correspondentes às operações a serem executadas (GetAll, Create, etc.) devem ser implementados.

import (
	"Exercicio1GoWebTarde/Internal/products/models"
)

type Service interface {
	GetAll() ([]models.Product, error)
	Create(product models.Product) (models.Product, error)
	GetByID(id string) (models.Product, error)
	Update(id string, product models.Product) (models.Product, error)
	Delete(id string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) GetAll() ([]models.Product, error) {
	return s.repo.GetAll()
}

func (s *service) Create(product models.Product) (models.Product, error) {
	return s.repo.Create(product)
}

func (s *service) GetByID(id string) (models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *service) Update(id string, product models.Product) (models.Product, error) {
	return s.repo.Update(id, product)
}

func (s *service) Delete(id string) error {
	return s.repo.Delete(id)
}
