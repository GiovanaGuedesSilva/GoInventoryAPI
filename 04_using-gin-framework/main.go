package main

/*
	Etapa 4 - Substituindo net/http por Gin

	Aqui introduzimos o framework Gin, uma das bibliotecas mais populares
	para criação de APIs em Go, que oferece roteamento mais poderoso,
	middlewares e tratamento simplificado de JSON.

	Esta etapa configura um servidor Gin e define uma rota GET que responde com JSON.
*/

import (
	"github.com/gin-gonic/gin" // Framework para criação de APIs HTTP em Go
)

func main() {
	// Cria um roteador com configurações padrão do Gin
	router := gin.Default()

	// Define uma rota GET na raiz "/" que será tratada pela função handler
	router.GET("/", handler)

	// Inicia o servidor na porta 8080
	router.Run(":8080")
}

// handler é a função que responde à rota GET "/"
func handler(c *gin.Context) {
	// Retorna um JSON com a mensagem "Hello World" e status 200 (OK)
	c.JSON(200, gin.H{"message": "Hello World"})
}
