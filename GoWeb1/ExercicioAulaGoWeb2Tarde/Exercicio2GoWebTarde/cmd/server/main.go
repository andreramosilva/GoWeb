package main

// O main do programa.
// O repositório, serviço e handler devem ser importados e injetados
// O roteador deve ser implementado para os diferentes endpoints

// O pacote handler com o controlador da entidade selecionada.
// A estrutura request deve ser gerada
// A estrutura do controller que tem o serviço como campo deve ser gerada
// A função que retorna o controlador deve ser gerada
// Todos os métodos correspondentes aos endpoints devem ser gerados

import (
	"Exercicio1GoWebTarde/Internal/products/handler"
	"Exercicio1GoWebTarde/Internal/products/repository"
	"Exercicio1GoWebTarde/Internal/products/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	repo := repository.NewRepository()
	serv := service.NewService(repo)
	handler.NewHandler(serv, r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
