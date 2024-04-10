package main

import (
	"encoding/json"
	"fmt"

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

func main() {
	jsonData := `[{"id":1,"name":"Product 1","color":"Red","price":10.5,"stock":100,"code":"ABC123","publish":true,"creationDate":"2021-01-01"},{"id":2,"name":"Product 2","color":"Blue","price":20.5,"stock":200,"code":"DEF456","publish":false,"creationDate":"2021-01-02"}]`

	var p []product
	err := json.Unmarshal([]byte(jsonData), &p)
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": p,
		})
	})
	router.Run(":8080")

}
