package main

/*
	Etapa 12 - Interface do Usecase (itemUsecasePort)

	Nesta etapa, introduzimos uma interface (itemUsecasePort) para a camada de usecase.
	O objetivo é desacoplar o handler da implementação concreta da lógica de negócio.
	Isso permite maior flexibilidade, facilita testes unitários com mocks
	e é um passo importante rumo à arquitetura limpa.

	A interface serve para desacoplar o código. Isso quer dizer:
		O handler não se importa como o usecase funciona por dentro.
		Ele só precisa saber: "saveItem salva um item" e "listItems lista os iten
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
	router := gin.Default()

	// Inicializa repositório e usecase
	r := newRepository()
	u := newItemUsecase(r)

	// Injeta a interface no handler (injeção por abstração)
	h := newHandler(u)

	// Rotas principais
	router.POST("/items", h.saveItem)
	router.GET("/items", h.listItems)

	log.Println("Server started at http://localhost:8080/")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

///////////////////////////////////////////////////////////////////////////////
// Erro global
///////////////////////////////////////////////////////////////////////////////

var errNotFound = errors.New("not found")

///////////////////////////////////////////////////////////////////////////////
// Handler
///////////////////////////////////////////////////////////////////////////////

type handler struct {
	usecase itemUsecasePort // interface do usecase (injeção por abstração)
}

func newHandler(u itemUsecasePort) *handler {
	return &handler{
		usecase: u,
	}
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
// Repository (armazenamento em memória)
///////////////////////////////////////////////////////////////////////////////

type repository struct {
	items mapRepo
}

func newRepository() *repository {
	return &repository{
		items: make(mapRepo),
	}
}

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

func (r *repository) listItems() (mapRepo, error) {
	return r.items, nil
}

///////////////////////////////////////////////////////////////////////////////
// Usecase (regra de negócio)
///////////////////////////////////////////////////////////////////////////////

// itemUsecasePort define as operações que o handler espera da camada de negócio
// O handler fala: "Me dá qualquer coisa que saiba salvar e listar item."
type itemUsecasePort interface {
	saveItem(item) error
	listItems() (mapRepo, error)
}

// Implementação concreta da interface
type itemUsecase struct {
	repo *repository
}

// Construtor retorna a interface, não a struct concreta
func newItemUsecase(repo *repository) itemUsecasePort {
	return &itemUsecase{
		repo: repo,
	}
}

func (u *itemUsecase) saveItem(it item) error {
	if err := u.repo.saveItem(it); err != nil {
		return fmt.Errorf("error saving item: %w", err)
	}
	return nil
}

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
// Domain (entidade de negócio)
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

// mapRepo é um alias para o tipo de armazenamento em memória
type mapRepo map[int]item

// Handler → Recebe a requisição HTTP e envia a resposta.
// Usecase (caso de uso) → Onde está a lógica de negócio (ex: validar se o item já existe).
// Repository → Armazena os dados (neste caso, na memória).
// Domain → Define o que é um "item" (a estrutura dos dados).
// Interface (itemUsecasePort) → Uma "promessa" do que a camada de negócio precisa saber fazer.
