package main

/*
	Etapa 4 - Substituição de net/http por Gin

	Nesta etapa, trocamos o uso da biblioteca padrão `net/http` pelo framework `Gin`,
	que facilita a criação de APIs RESTful com mais flexibilidade e menos código repetitivo.

	As rotas implementadas são equivalentes à etapa anterior,
	mas agora com handlers separados para cada método HTTP usando recursos do Gin.
*/

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa o roteador com configurações padrão (logger e recovery middleware)
	router := gin.Default()

	// Define as rotas para a URL "/"
	router.GET("/", homeGet)
	router.POST("/", homePost)
	router.PUT("/", homePut)
	router.DELETE("/", homeDelete)

	// Outras rotas
	router.GET("/hello", hello)
	router.POST("/bye", bye)

	// Inicia o servidor na porta 8080
	router.Run(":8080")
}

// Handlers para a rota "/"
func homeGet(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the home page!")
}

func homePost(c *gin.Context) {
	c.String(http.StatusOK, "Post to the home page!")
}

func homePut(c *gin.Context) {
	c.String(http.StatusOK, "Put to the home page!")
}

func homeDelete(c *gin.Context) {
	c.String(http.StatusOK, "Delete the home page!")
}

// Handler para a rota "/hello"
func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

// Handler para a rota "/bye"
func bye(c *gin.Context) {
	c.String(http.StatusOK, "Goodbye guys!")
}
