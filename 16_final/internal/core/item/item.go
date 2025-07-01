package item

import (
	"time"
)

/*
Item representa um produto ou item do sistema.

Cada campo possui uma tag `json` indicando como será serializado/deserializado
quando os dados forem enviados ou recebidos via API REST.
*/
type Item struct {
	ID          int       `json:"id"`          // Identificador único do item
	Code        string    `json:"code"`        // Código interno ou SKU
	Title       string    `json:"title"`       // Nome ou título do item
	Description string    `json:"description"` // Descrição detalhada
	Price       float64   `json:"price"`       // Preço do item
	Stock       int       `json:"stock"`       // Quantidade disponível em estoque
	Status      string    `json:"status"`      // Status do item (ex: ativo, inativo, esgotado)
	CreatedAt   time.Time `json:"created_at"`  // Data de criação do item
	UpdatedAt   time.Time `json:"updated_at"`  // Última data de atualização
}

/*
MapRepo representa uma estrutura de dados do tipo mapa
usada na implementação de repositório em memória.

Chave: ID do item (int)
Valor: struct Item
*/
type MapRepo map[int]Item
