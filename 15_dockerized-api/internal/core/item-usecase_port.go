package core

import "api/internal/core/item"

/*
	ItemUsecasePort define a interface para o caso de uso relacionado a itens.

	Essa interface representa as ações que a camada de aplicação oferece ao mundo externo (ex: handlers HTTP).
	Ela serve como contrato entre os adaptadores (como os handlers REST) e a lógica de negócio.
*/
type ItemUsecasePort interface {
	// SaveItem salva um novo item no sistema (pode ser no banco ou na memória, dependendo do repositório injetado)
	SaveItem(item.Item) error

	// ListItems retorna todos os itens existentes no sistema
	ListItems() (item.MapRepo, error)
}
