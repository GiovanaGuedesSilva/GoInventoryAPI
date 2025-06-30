package main

/*
	Etapa 5 - Explicação Detalhada da Estrutura com Gin

	Nesta etapa, mantemos o mesmo código funcional da Etapa 4, porém com explicações mais detalhadas
	sobre a estrutura do Gin e o que está acontecendo em cada parte da aplicação.

	Nosso objetivo aqui é reforçar o entendimento sobre como o framework Gin trabalha com rotas e handlers,
	e como isso se compara com o que fizemos anteriormente com net/http.
*/

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Cria uma instância do roteador padrão do Gin.
	// Ele vem com middleware embutido de logger e recuperação de panics.
	router := gin.Default()

	// ROTAS PARA "/"
	// Cada método HTTP (GET, POST, PUT, DELETE) é mapeado para um handler específico.
	// Isso facilita a leitura e manutenção do código, além de seguir boas práticas REST.

	router.GET("/", homeGet)       // Requisições GET à rota "/"
	router.POST("/", homePost)     // Requisições POST à rota "/"
	router.PUT("/", homePut)       // Requisições PUT à rota "/"
	router.DELETE("/", homeDelete) // Requisições DELETE à rota "/"

	// OUTRAS ROTAS
	// Aqui mostramos que o mesmo roteador pode lidar com diferentes caminhos e métodos.

	router.GET("/hello", hello) // Responde a um GET com saudação
	router.POST("/bye", bye)    // Responde a um POST com despedida

	// Inicia o servidor web na porta 8080
	router.Run(":8080")
}

// --- HANDLERS ---
// Cada handler recebe um ponteiro para o contexto (c *gin.Context),
// que oferece acesso facilitado à requisição, resposta, parâmetros, JSON, status HTTP, etc.

// homeGet lida com GET /
func homeGet(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the home page!")
}

// homePost lida com POST /
func homePost(c *gin.Context) {
	c.String(http.StatusOK, "Post to the home page!")
}

// homePut lida com PUT /
func homePut(c *gin.Context) {
	c.String(http.StatusOK, "Put to the home page!")
}

// homeDelete lida com DELETE /
func homeDelete(c *gin.Context) {
	c.String(http.StatusOK, "Delete the home page!")
}

// hello lida com GET /hello
func hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

// bye lida com POST /bye
func bye(c *gin.Context) {
	c.String(http.StatusOK, "Goodbye guys!")
}
