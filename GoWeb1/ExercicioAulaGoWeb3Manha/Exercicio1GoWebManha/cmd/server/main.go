package main

import (
	"Exercicio1GoWebManha/Internal/database"
	"Exercicio1GoWebManha/Internal/middleware"
	"Exercicio1GoWebManha/Internal/products/handler"
	"Exercicio1GoWebManha/Internal/products/repository"
	"Exercicio1GoWebManha/Internal/products/service"
	"Exercicio1GoWebManha/Internal/server"
	"Exercicio1GoWebManha/Internal/transactions/handler"
	"Exercicio1GoWebManha/Internal/transactions/repository"
	"Exercicio1GoWebManha/Internal/transactions/service"
	"Exercicio1GoWebManha/Internal/users/handler"
	"Exercicio1GoWebManha/Internal/users/repository"
	"Exercicio1GoWebManha/Internal/users/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.Connect()
	r := mux.NewRouter()
	// Products
	productsRepo := repository.NewRepository()
	productsService := service.NewService(productsRepo)
	productsHandler := handler.NewHandler(productsService)
	productsHandler.Register(r)
	// Users
	usersRepo := repository.NewRepository()
	usersService := service.NewService(usersRepo)
	usersHandler := handler.NewHandler(usersService)
	usersHandler.Register(r)
	// Transactions
	transactionsRepo := repository.NewRepository()
	transactionsService := service.NewService(transactionsRepo)
	transactionsHandler := handler.NewHandler(transactionsService)
	transactionsHandler.Register(r)
	// Middleware
	r.Use(middleware.Logger)
	// Server
	srv := server.NewServer(r)
	log.Fatal(http.ListenAndServe(":8080", srv))
}
