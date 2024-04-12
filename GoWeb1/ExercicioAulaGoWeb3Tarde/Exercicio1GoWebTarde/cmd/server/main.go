package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	usuario := os.Getenv("MY_USER")
	senha := os.Getenv("MY_PASSWORD")

	log.Println("Usuario: ", usuario)
	log.Println("Senha: ", senha)

}
