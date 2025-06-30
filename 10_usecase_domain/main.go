package main

/*
	Etapa 10 - Separação por Domínio e Usecase

	Nesta etapa, a API começa a adotar uma estrutura mais limpa e escalável,
	introduzindo a separação de camadas:
	- Domain: definição da entidade (item)
	- Usecase: lógica de negócio central
	- Handler: interface com o mundo externo (HTTP)

	Essa arquitetura facilita manutenção, testes e futura persistência de dados.
*/

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Cria o usecase e injeta no handler
	u := newItemUsecase()
	h := newHandler(u)

	// Inicializa o roteador do Gin
	router := gin.Default()

	// Rotas básicas
	router.GET("/", h.home)
	router.GET("/hello", h.hello)
	router.POST("/bye", h.bye)

	// Novas rotas de item
	router.POST("/items", h.saveItem)
	router.GET("/items", h.listItems)

	// Log e inicialização do servidor
	log.Println("Server started at http://localhost:8080/")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

///////////////////////////////////////////////////////////////////////////////
// Global error
///////////////////////////////////////////////////////////////////////////////

// Erro de item não encontrado (simulado)
var errNotFound = errors.New("not found")

///////////////////////////////////////////////////////////////////////////////
// Handler - Interface HTTP
///////////////////////////////////////////////////////////////////////////////

type handler struct {
	usecase *itemUsecase
}

// Construtor do handler com injeção de dependência
func newHandler(u *itemUsecase) *handler {
	return &handler{usecase: u}
}

// Handler GET /
func (h *handler) home(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the home page!")
}

// Handler GET /hello
func (h *handler) hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

// Handler POST /bye - recebe JSON simples
func (h *handler) bye(c *gin.Context) {
	var msg map[string]string
	if err := c.BindJSON(&msg); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}
	message, exists := msg["message"]
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Message field is missing"})
		return
	}
	c.String(http.StatusOK, "Received POST request with message: %s", message)
}

// Handler POST /items - salva item enviado via JSON
func (h *handler) saveItem(c *gin.Context) {
	var it item

	// Tenta fazer o binding do JSON recebido para a struct item
	if err := c.BindJSON(&it); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chama o usecase para salvar
	if err := h.usecase.saveItem(it); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "item saved successfully")
}

// Handler GET /items - lista todos os itens (mock)
func (h *handler) listItems(c *gin.Context) {
	its, err := h.usecase.listItems()
	if err != nil {
		if err == errNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, its)
}

///////////////////////////////////////////////////////////////////////////////
// Usecases - Regra de negócio
///////////////////////////////////////////////////////////////////////////////

type itemUsecase struct{}

// Construtor
func newItemUsecase() *itemUsecase {
	return &itemUsecase{}
}

// Salva o item (simulado - ainda não persiste de verdade)
func (u *itemUsecase) saveItem(it item) error {
	// Aqui será implementada a lógica de persistência
	_ = it // Simulação
	return nil
}

// Lista os itens (simulado - retorno fixo e vazio)
func (u *itemUsecase) listItems() (map[int]item, error) {
	// Aqui será implementada a lógica real de busca no repositório
	its := make(map[int]item) // retorno mock
	return its, nil
}

///////////////////////////////////////////////////////////////////////////////
// Domain - Entidade item
///////////////////////////////////////////////////////////////////////////////

// item representa um produto do inventário
type item struct {
	ID          int       `json:"id"`          // ID único
	Code        string    `json:"code"`        // Código do produto
	Title       string    `json:"title"`       // Nome/título
	Description string    `json:"description"` // Descrição detalhada
	Price       float64   `json:"price"`       // Preço
	Stock       int       `json:"stock"`       // Quantidade em estoque
	Status      string    `json:"status"`      // Status (ativo, inativo etc.)
	CreatedAt   time.Time `json:"created_at"`  // Data de criação
	UpdatedAt   time.Time `json:"updated_at"`  // Data de atualização
}
