package handler

// O pacote handler com o controlador da entidade selecionada.
// A estrutura request deve ser gerada
// A estrutura do controller que tem o serviço como campo deve ser gerada
// A função que retorna o controlador deve ser gerada
// Todos os métodos correspondentes aos endpoints devem ser gerados

import (
	"Exercicio1GoWebManha/Internal/products"
	"Exercicio1GoWebManha/Internal/products/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type request struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type controller struct {
	service products.Service
}

func NewController(service products.Service) *controller {
	return &controller{service}
}

func (c *controller) GetAll(ctx echo.Context) error {
	products, err := c.service.GetAll()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, products)
}

func (c *controller) Create(ctx echo.Context) error {
	var req request
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	product, err := c.service.Create(models.Product{
		Name:  req.Name,
		Price: req.Price,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, product)
}

func (c *controller) GetByID(ctx echo.Context) error {
	id := ctx.Param("id")
	product, err := c.service.GetByID(id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err.Error())
	}
	return ctx.JSON(http.StatusOK, product)
}

func (c *controller) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	var req request
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}
	product, err := c.service.Update(id, models.Product{
		Name:  req.Name,
		Price: req.Price,
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, product)
}

func (c *controller) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	err := c.service.Delete(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusNoContent)
}
