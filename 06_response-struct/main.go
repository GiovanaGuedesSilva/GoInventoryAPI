package main

/*
	Etapa 6 - Criando uma Struct para Resposta JSON

	Nesta etapa, substituímos o uso de `gin.H` (mapa genérico) por uma struct nomeada.
	Essa abordagem torna o código mais seguro, legível e preparado para documentação automática,
	como Swagger ou validação de respostas.
*/

import (
	"github.com/gin-gonic/gin"
)

// MessageResponse define a estrutura da resposta JSON
type MessageResponse struct {
	Message string `json:"message"` // Campo "message" aparecerá como chave no JSON
}

func main() {
	// Inicializa o roteador padrão do Gin
	router := gin.Default()

	// Define a rota GET para a raiz "/"
	router.GET("/", handler)

	// Inicia o servidor na porta 8080
	router.Run(":8080")
}

// handler lida com requisições GET e retorna um JSON baseado na struct MessageResponse
func handler(c *gin.Context) {
	// Cria uma instância da resposta
	response := MessageResponse{
		Message: "Hello World",
	}

	// Envia a resposta como JSON com status 200 OK
	c.JSON(200, response)
}
