package item

/*
ItemRepositoryPort define o contrato que toda implementação de repositório de itens deve cumprir.

Essa interface serve como uma "porta" entre a lógica de negócio (usecases) e os detalhes de persistência
(banco de dados, memória, etc).

Vantagens:
----------
- A lógica de negócio depende da abstração, não da implementação;
- Permite trocar facilmente o repositório (MySQL, MongoDB, memória, etc.);
- Facilita testes, pois podemos usar mocks ou versões em memória;
- Segue os princípios da arquitetura limpa (clean architecture).
*/
type ItemRepositoryPort interface {
	// SaveItem salva um novo item no repositório.
	// Retorna erro caso o item viole regras (ex: ID duplicado).
	SaveItem(*Item) error

	// ListItems retorna todos os itens armazenados no repositório.
	// Retorna o mapa de itens e um erro (se houver).
	ListItems() (MapRepo, error)

	// UpdateItem atualiza um item existente no repositório.
	// Retorna erro caso o item não exista.
	UpdateItem(*Item) error

	// DeleteItem remove um item com base no ID.
	// Retorna erro caso o item não exista ou ID seja inválido.
	DeleteItem(int) error
}
