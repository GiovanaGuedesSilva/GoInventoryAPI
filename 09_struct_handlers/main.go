package main

/*
	Etapa 9 - Uso de Struct para Organizar Handlers

	Nesta etapa, os handlers da API foram encapsulados em uma struct `handler`.

	Isso permite organizar melhor o código, facilitar testes e possibilitar a injeção de dependências
	como bancos de dados, serviços ou configurações em projetos mais complexos.

	Os métodos da struct `handler` substituem as funções soltas das etapas anteriores.
*/

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Cria uma nova instância do handler
	h := newHandler()

	// Inicializa o roteador do Gin
	router := gin.Default()

	// Define as rotas e associa aos métodos da struct
	router.GET("/", h.home)
	router.GET("/hello", h.hello)
	router.POST("/bye", h.bye)

	// Log informativo
	log.Println("Server started at http://localhost:8080/")

	// Inicia o servidor e trata erro
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

// handler define uma estrutura que agrupa os métodos de tratamento das rotas
type handler struct{}

// newHandler retorna uma nova instância de handler
func newHandler() *handler {
	return &handler{}
}

// Método que trata GET /
func (h *handler) home(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the home page!")
}

// Método que trata GET /hello
func (h *handler) hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

// Método que trata POST /bye com entrada em JSON
func (h *handler) bye(c *gin.Context) {
	var msg map[string]string

	// Tenta fazer o binding do JSON
	if err := c.BindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	// Verifica se a chave "message" está presente
	message, exists := msg["message"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Message field is missing"})
		return
	}

	// Envia a resposta com a mensagem recebida
	c.String(http.StatusOK, "Received POST request with message: %s", message)
}
