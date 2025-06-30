package main

/*
	Etapa 8:
	Objetivo: Receber e processar dados em JSON via POST.
	Agora o handler do endpoint "/bye" usa o método BindJSON do Gin para decodificar automaticamente um JSON no corpo da requisição.
	Isso representa uma prática comum em APIs REST reais.

	Exemplo de JSON esperado:
	{
		"message": "Olá, Giovana!"
	}
*/

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", home)
	router.GET("/hello", hello)
	router.POST("/bye", bye)

	// Logs no terminal (visíveis quando o servidor inicia)
	fmt.Println("Welcome!")
	log.Println("Server started at http://localhost:8080/")

	// Inicia o servidor na porta 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// Rota GET "/"
func home(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the home page!")
}

// Rota GET "/hello"
func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

// Rota POST "/bye" recebendo JSON {"message": "alguma coisa"}
func bye(c *gin.Context) {
	var msg map[string]string

	// Tenta ler o JSON para dentro do map
	if err := c.BindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	// Verifica se o campo "message" existe
	message, exists := msg["message"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Message field is missing"})
		return
	}

	c.String(http.StatusOK, "Received POST request with message: %s", message)
}
