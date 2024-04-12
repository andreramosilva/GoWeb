// A funcionalidade para salvar as informações da solicitação em um arquivo json deve ser implementada, para isso devem ser realizadas as seguintes etapas:

// Ao invés de salvar os valores da nossa entidade na memória, deve ser criado um arquivo; os valores que são adicionados são salvos nele.

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

func CreateProduct(c *gin.Context) {
	var prod product
	err := c.BindJSON(&prod)
	if err != nil {
		fmt.Println(err)
	}

	jsonData, _ := os.ReadFile("../../products.json")

	var products []product
	err = json.Unmarshal([]byte(jsonData), &products)
	if err != nil {
		fmt.Println(err)
	}

	products = append(products, prod)

	file, _ := json.MarshalIndent(products, "", " ")
	_ = os.WriteFile("../../products.json", file, 0644)

	c.JSON(201, prod)
}

func SearchProduct(c *gin.Context) {
	id := c.Param("id")
	jsonData, _ := os.ReadFile("../../products.json")

	var products []product
	err := json.Unmarshal([]byte(jsonData), &products)
	if err != nil {
		fmt.Println(err)
	}

	for _, prod := range products {
		y, _ := strconv.Atoi(id)
		if prod.Id == y {
			c.JSON(200, prod)
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
	server.POST("/produtos", CreateProduct)
	server.Run(":8080")

}
