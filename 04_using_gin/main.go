package main

/*
	Etapa 4:
	Objetivo: Substituir o servidor HTTP nativo pelo framework Gin.
	O Gin facilita o roteamento e tratamento de requisições HTTP, além de fornecer ferramentas mais robustas para desenvolvimento de APIs.
	Aqui, definimos handlers para diferentes métodos HTTP (GET, POST, PUT, DELETE) na rota "/" e também para outras rotas (/hello e /bye).
*/

import (
	"net/http"

	"github.com/gin-gonic/gin" // Importa o framework Gin
)

func main() {
	// Cria um roteador Gin com middlewares padrão (logger, recovery, etc.)
	router := gin.Default()

	// Middlewares são funções que interceptam as requisições e/ou respostas entre o cliente e o servidor.
	// Eles funcionam como um "meio do caminho" (daí o nome middleware) entre o que chega à sua API e o que ela vai responder.
	// Eles são usados para executar lógica comum antes ou depois do handler principal da rota.

	// Define handlers para a rota "/" com diferentes métodos HTTP
	router.GET("/", homeGet)
	router.POST("/", homePost)
	router.PUT("/", homePut)
	router.DELETE("/", homeDelete)

	// Define rota "/hello" com método GET
	router.GET("/hello", hello)

	// Define rota "/bye" com método POST
	router.POST("/bye", bye)

	// Inicia o servidor HTTP na porta 8080
	router.Run(":8080")
}

// Handler para GET "/"
func homeGet(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the home page!")
}

// Handler para POST "/"
func homePost(c *gin.Context) {
	c.String(http.StatusOK, "Post to the home page!")
}

// Handler para PUT "/"
func homePut(c *gin.Context) {
	c.String(http.StatusOK, "Put to the home page!")
}

// Handler para DELETE "/"
func homeDelete(c *gin.Context) {
	c.String(http.StatusOK, "Delete the home page!")
}

// Handler para GET "/hello"
func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

// Handler para POST "/bye"
func bye(c *gin.Context) {
	c.String(http.StatusOK, "Goodbye guys!")
}
