package main

/*
	Etapa 11 - Adição da Camada de Repositório

	Nesta etapa, foi introduzida a camada de persistência simulada (`repository`),
	que armazena os dados em memória utilizando um `map[int]item`.

	Agora a arquitetura possui:
	- Handler (interface HTTP)
	- Usecase (lógica de negócio)
	- Repository (persistência)
	- Domain (entidade `item`)

	Isso aproxima a API de uma arquitetura limpa e escalável.
*/

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicia o repositório (persistência em memória)
	r := newRepository()

	// Injeta o repositório no usecase
	u := newItemUsecase(r)

	// Injeta o usecase no handler
	h := newHandler(u)

	// Cria roteador Gin
	router := gin.Default()

	// Rotas de teste
	router.GET("/", h.home)
	router.GET("/hello", h.hello)
	router.POST("/bye", h.bye)

	// Rotas de item
	router.POST("/items", h.saveItem)
	router.GET("/items", h.listItems)

	log.Println("Server started at http://localhost:8080/")

	// Inicia o servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

///////////////////////////////////////////////////////////////////////////////
// Erros globais
///////////////////////////////////////////////////////////////////////////////

var errNotFound = errors.New("not found")

///////////////////////////////////////////////////////////////////////////////
// Camada Handler (interface HTTP)
///////////////////////////////////////////////////////////////////////////////

type handler struct {
	usecase *itemUsecase
}

// Construtor do handler
func newHandler(u *itemUsecase) *handler {
	return &handler{usecase: u}
}

// Rota GET /
func (h *handler) home(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the home page!")
}

// Rota GET /hello
func (h *handler) hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello, world!")
}

// Rota POST /bye
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

// Rota POST /items
func (h *handler) saveItem(c *gin.Context) {
	var it item
	err := c.BindJSON(&it)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.saveItem(it); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "item saved successfully")
}

// Rota GET /items
func (h *handler) listItems(c *gin.Context) {
	items, err := h.usecase.listItems()
	if err != nil {
		if err == errNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, items)
}

///////////////////////////////////////////////////////////////////////////////
// Camada Usecase (regras de negócio)
///////////////////////////////////////////////////////////////////////////////

type itemUsecase struct {
	repo *repository
}

// Construtor do usecase
func newItemUsecase(repo *repository) *itemUsecase {
	return &itemUsecase{repo: repo}
}

// Lógica para salvar item (com validação de ID e duplicidade)
func (u *itemUsecase) saveItem(it item) error {
	if err := u.repo.saveItem(it); err != nil {
		return fmt.Errorf("error saving item: %w", err)
	}
	return nil
}

// Lógica para listar todos os itens
func (u *itemUsecase) listItems() (mapRepo, error) {
	its, err := u.repo.listItems()
	if err != nil {
		return nil, fmt.Errorf("error in repository: %w", err)
	}
	if len(its) == 0 {
		return nil, errNotFound
	}
	return its, nil
}

///////////////////////////////////////////////////////////////////////////////
// Camada Repository (persistência em memória)
///////////////////////////////////////////////////////////////////////////////

// Estrutura de repositório com um mapa
type repository struct {
	items mapRepo
}

// Construtor
func newRepository() *repository {
	return &repository{items: make(mapRepo)}
}

// Salva item no mapa
func (r *repository) saveItem(it item) error {
	if it.ID == 0 {
		return fmt.Errorf("item ID cannot be 0")
	}
	if _, exists := r.items[it.ID]; exists {
		return fmt.Errorf("item with ID %d already exists", it.ID)
	}
	r.items[it.ID] = it
	return nil
}

// Retorna todos os itens
func (r *repository) listItems() (mapRepo, error) {
	return r.items, nil
}

///////////////////////////////////////////////////////////////////////////////
// Camada Domain (entidade de negócio)
///////////////////////////////////////////////////////////////////////////////

// item representa um produto do inventário
type item struct {
	ID          int       `json:"id"`
	Code        string    `json:"code"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Alias para mapa de itens
type mapRepo map[int]item
