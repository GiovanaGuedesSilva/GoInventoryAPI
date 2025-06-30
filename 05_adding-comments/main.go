package main

/*
	Etapa 5 - Comentando o Código com Gin

	Nesta etapa, não há mudança funcional em relação à etapa anterior.
	O objetivo é adicionar comentários explicativos no código para fins didáticos,
	tornando mais claro o funcionamento da API com o framework Gin.
*/

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Cria um roteador padrão do Gin com middleware de logging e recovery
	router := gin.Default()

	// Define uma rota GET na raiz "/"
	router.GET("/", handler)

	// Inicia o servidor e escuta na porta 8080
	router.Run(":8080")
}

// handler lida com requisições GET para a rota "/"
func handler(c *gin.Context) {
	// Retorna uma resposta JSON com status 200 OK
	c.JSON(200, gin.H{
		"message": "Hello World", // Conteúdo da resposta
	})
}
