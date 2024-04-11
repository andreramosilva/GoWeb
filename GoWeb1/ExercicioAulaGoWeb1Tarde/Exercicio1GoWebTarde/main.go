// Exercício 1 - Vamos filtrar nosso endpoint
// Dependendo do tema escolhido, precisamos adicionar filtros ao nosso endpoint, ele deve ser capaz de filtrar todos os campos.

// Dentro do manipulador de endpoint, recebi os valores para filtrar do contexto.
// Em seguida, ele gera a lógica do filtro para nossa matriz.
// Retorne a matriz filtrada por meio do endpoint.

package main

import (
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
	c.JSON(200, gin.H{
		"message": "Produto encontrado",
		"id":      id,
	})
}

func main() {
	// jsonData := `[{"id":1,"name":"Product 1","color":"Red","price":10.5,"stock":100,"code":"ABC123","publish":true,"creationDate":"2021-01-01"},{"id":2,"name":"Product 2","color":"Blue","price":20.5,"stock":200,"code":"DEF456","publish":false,"creationDate":"2021-01-02"}]`

	// var p []product
	// err := json.Unmarshal([]byte(jsonData), &p)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// router := gin.Default()
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"data": p,
	// 	})
	// })
	// router.Run(":8080")

	server := gin.Default()
	server.GET("/", RootPage)
	server.GET("/produtos/:id", SearchProduct)
	server.Run(":8080")

}
