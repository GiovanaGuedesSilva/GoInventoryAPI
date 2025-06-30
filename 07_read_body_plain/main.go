package main

/*
	Etapa 7 - Leitura de Dados do Corpo da Requisição (POST)

	Nesta etapa, começamos a tratar dados vindos do cliente via corpo da requisição HTTP (body).

	Usamos `io.ReadAll(c.Request.Body)` para ler o conteúdo bruto (raw).
	Isso é útil quando o cliente envia texto puro (plain text), e marca o primeiro passo rumo ao consumo de JSON.
*/

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa o roteador do Gin
	router := gin.Default()

	// Rotas
	router.GET("/", home)
	router.GET("/hello", hello)
	router.POST("/bye", bye)

	// Inicia o servidor
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

// bye responde a POST /bye, lendo o corpo da requisição como texto puro
func bye(c *gin.Context) {
	// Tenta ler todo o corpo da requisição
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		// Retorna erro se a leitura falhar
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	// Converte o conteúdo do body para string
	message := string(body)

	// Responde com a mensagem recebida
	c.String(http.StatusOK, "Received POST request with message: %s", message)
}
