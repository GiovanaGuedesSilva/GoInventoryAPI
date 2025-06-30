package main

/*
	Etapa 7:
	Objetivo: Ler o corpo da requisição (request body) manualmente.
	Aqui, ao enviar uma requisição POST para "/bye", o servidor lê o conteúdo bruto enviado e o exibe como resposta.
	Isso ensina como acessar o corpo de uma requisição sem depender de JSON ou estruturas.
*/

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", home)
	router.GET("/hello", hello)
	router.POST("/bye", bye)

	router.Run(":8080")
}

// Rota GET "/"
func home(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the home page!")
}

// Rota GET "/hello"
func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

// Rota POST "/bye" lendo o corpo da requisição manualmente
func bye(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}
	message := string(body)
	c.String(http.StatusOK, "Received POST request with message: %s", message)
}
