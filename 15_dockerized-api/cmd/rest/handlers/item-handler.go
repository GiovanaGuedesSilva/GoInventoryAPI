package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api/internal/core"
	"api/internal/core/item"
	"api/pkg/config"
)

/*
handler é o responsável por lidar com as requisições HTTP relacionadas aos "itens".
Ele atua como um intermediário entre a camada de apresentação (HTTP) e a camada de lógica de negócio (caso de uso).
*/
type handler struct {
	core core.ItemUsecasePort // Interface da camada de caso de uso relacionada a "item"
}

/*
NewHandler cria uma nova instância de handler, recebendo como dependência
um objeto que implementa a interface ItemUsecasePort.
Isso permite injeção de dependência e facilita testes.
*/
func NewHandler(u core.ItemUsecasePort) *handler {
	return &handler{
		core: u,
	}
}

/*
SaveItem lida com a requisição HTTP para salvar um novo item.

Passos:
 1. Tenta fazer o bind do JSON recebido no corpo da requisição para a struct `item.Item`.
    Se falhar (ex: JSON inválido), retorna erro 400 (Bad Request).
 2. Chama o método da camada de caso de uso `SaveItem` passando o item.
    Se ocorrer erro ao salvar (ex: problema no banco), retorna erro 500 (Internal Server Error).
 3. Se tudo correr bem, retorna status 200 com mensagem de sucesso.
*/
func (h *handler) SaveItem(c *gin.Context) {
	var it item.Item

	// Tenta converter o JSON recebido para a struct `Item`
	err := c.BindJSON(&it)
	if err != nil {
		// Se falhar, retorna erro de requisição inválida
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Chama o caso de uso para salvar o item
	if err := h.core.SaveItem(it); err != nil {
		// Se falhar, retorna erro interno do servidor
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Se sucesso, retorna status 200
	c.JSON(http.StatusOK, "item salvo com sucesso")
}

/*
ListItems lida com a requisição HTTP para listar todos os itens cadastrados.

Passos:
1. Chama o método da camada de caso de uso `ListItems`.
2. Se retornar erro "not found", retorna 404 (Not Found).
3. Se for outro erro, retorna 500 (Internal Server Error).
4. Se sucesso, retorna status 200 com a lista de itens no corpo da resposta.
*/
func (h *handler) ListItems(c *gin.Context) {
	// Busca todos os itens
	its, err := h.core.ListItems()
	if err != nil {
		// Verifica se o erro é "não encontrado"
		if err == config.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			// Se for outro erro qualquer
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Retorna a lista de itens com status 200
	c.JSON(http.StatusOK, its)
}
