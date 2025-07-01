package core

import "api/internal/core/item"

/*
ItemUsecasePort define a interface da camada de aplicação para a entidade Item.

Essa interface é usada para desacoplar a lógica de negócio da camada de entrega (handlers REST, gRPC, etc).

Ela descreve as ações que podem ser realizadas com itens, sem se preocupar com os detalhes de onde os dados vêm ou vão.

Vantagens:
----------
- Permite usar diferentes implementações do usecase (ex: para testes);
- Permite que os handlers dependam da abstração, não da implementação;
- Mantém a aplicação desacoplada, coesa e testável.
*/
type ItemUsecasePort interface {
	// SaveItem salva um novo item, validando e repassando para o repositório.
	SaveItem(item.Item) error

	// ListItems retorna todos os itens cadastrados.
	ListItems() (item.MapRepo, error)

	// UpdateItem atualiza os dados de um item existente.
	UpdateItem(item.Item) error

	// DeleteItem remove um item com base no seu ID.
	DeleteItem(int) error
}
