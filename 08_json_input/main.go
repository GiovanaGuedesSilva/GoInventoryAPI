package main

/*
	Etapa 8 - Leitura de JSON com Gin

	Nesta etapa, avançamos para a leitura de dados em formato JSON no corpo da requisição.

	Ao invés de ler o body como texto (etapa 7), usamos o método BindJSON do Gin
	para converter diretamente o JSON enviado em um mapa Go (`map[string]string`).

	Isso é mais seguro, limpo e escalável, especialmente em APIs RESTful modernas.
*/

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa o roteador
	router := gin.Default()

	// Rotas definidas
	router.GET("/", home)
	router.GET("/hello", hello)
	router.POST("/bye", bye)

	// Mensagens informativas no terminal
	fmt.Println("Welcome!")
	log.Println("Server started at http://localhost:8080/")

	// Inicia o servidor e trata erro de execução
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// home responde a GET /
func home(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the home page!")
}

// hello responde a GET /hello
func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

// bye responde a POST /bye, lendo JSON do corpo da requisição
func bye(c *gin.Context) {
	// Declaração de um mapa para armazenar os dados recebidos
	var msg map[string]string

	// Tenta fazer o binding do JSON recebido
	if err := c.BindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	// Verifica se o campo "message" existe no JSON
	message, exists := msg["message"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Message field is missing"})
		return
	}

	// Envia a resposta com a mensagem recebida
	c.String(http.StatusOK, "Received POST request with message: %s", message)
}
