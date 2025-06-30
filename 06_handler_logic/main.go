package main

/*
	Etapa 6 - Introdução de Lógica dentro dos Handlers

	Nesta etapa, mantemos a estrutura do Gin, mas iniciamos o uso de lógica dentro dos handlers.

	Na rota POST "/bye", simulamos o uso de uma variável local para formar uma resposta dinâmica.
	Esse é o primeiro passo para começar a processar dados enviados em requisições, como corpo em JSON ou formulários.
*/

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa o roteador padrão do Gin
	router := gin.Default()

	// Rotas definidas com handlers básicos
	router.GET("/", home)
	router.GET("/hello", hello)
	router.POST("/bye", bye)

	// Inicia o servidor na porta 8080
	router.Run(":8080")
}

// home responde a GET /
func home(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the home page!")
}

// hello responde a GET /hello
func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

// bye responde a POST /bye com lógica interna simulando o processamento de dados
func bye(c *gin.Context) {
	// Simula o uso de uma variável dinâmica
	message := "Hello and bye!"
	c.String(http.StatusOK, "Received POST request with message: %s", message)
}
