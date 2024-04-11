// Exercício 1 - Vamos filtrar nosso endpoint
// precisamos adicionar filtros ao nosso endpoint, ele deve ser capaz de filtrar todos os campos.

// Dentro do manipulador de endpoint, recebi os valores para filtrar do contexto.
// Em seguida, ele gera a lógica do filtro para nossa matriz.
// Retorne a matriz filtrada por meio do endpoint.

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	Color        string  `json:"color"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
	Code         string  `json:"code"`
	Publish      bool    `json:"publish"`
	CreationDate string  `json:"creationDate"`
}

func RootPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, Andre",
	})
}

func SearchProduct(c *gin.Context) {
	id := c.Param("id")
	file, err := os.ReadFile("../../products.json")
	if err != nil {
		fmt.Println(err)
	}
	var products []product
	err = json.Unmarshal(file, &products)
	if err != nil {
		fmt.Println(err)
	}

	for _, product := range products {
		y, _ := strconv.Atoi(id)
		print(y)
		if product.Id == y {
			c.JSON(200, product)
			return
		}
	}
	c.JSON(404, gin.H{
		"message": "Product not found",
	})

}

func main() {

	server := gin.Default()
	server.GET("/", RootPage)
	server.GET("/produtos/:id", SearchProduct)
	server.Run(":8080")

}
