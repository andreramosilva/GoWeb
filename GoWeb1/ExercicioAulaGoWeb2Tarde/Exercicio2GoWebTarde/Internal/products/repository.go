package products

// Repository, você deve ter acesso à variável armazenada na memória.
// O arquivo repository.go deve ser criado
// A estrutura da entidade deve ser criada
// Variáveis ​​globais devem ser criadas para armazenar as entidades
// A interface Repository deve ser gerada com todos os seus métodos
// A estrutura do repositório deve ser gerada
// Deve ser gerada uma função que retorne o Repositório
// Todos os métodos correspondentes às operações a serem executadas (GetAll, Store, etc.) devem ser implementados.

import (
	"Exercicio1GoWebTarde/Internal/products/models"
)

type Repository interface {
	GetAll() ([]models.Product, error)
	Create(product models.Product) (models.Product, error)
	GetByID(id string) (models.Product, error)
	Update(id string, product models.Product) (models.Product, error)
	Delete(id string) error
}

type repository struct {
	products []models.Product
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]models.Product, error) {
	return r.products, nil
}

func (r *repository) Create(product models.Product) (models.Product, error) {
	r.products = append(r.products, product)
	return product, nil
}

func (r *repository) GetByID(id string) (models.Product, error) {
	for _, product := range r.products {
		if product.ID == id {
			return product, nil
		}
	}
	return models.Product{}, nil
}

func (r *repository) Update(id string, product models.Product) (models.Product, error) {

	for i, p := range r.products {
		if p.ID == id {
			r.products[i] = product
			return product, nil
		}
	}
	return models.Product{}, nil
}

func (r *repository) Delete(id string) error {
	for i, product := range r.products {
		if product.ID == id {
			r.products = append(r.products[:i], r.products[i+1:]...)
			return nil
		}
	}
	return nil
}
