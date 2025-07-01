package item

/*
	ItemRepositoryPort define a interface (contrato) que todo repositório de itens deve seguir.

	Essa interface é usada pela camada de caso de uso, que depende dela — e **não de implementações concretas**.

	Isso permite que possamos trocar a implementação do repositório (por exemplo, de MySQL para memória)
	sem mudar a lógica de negócio (clean architecture).
*/
type ItemRepositoryPort interface {
	// SaveItem salva um item no repositório (MySQL, memória, etc.)
	SaveItem(*Item) error

	// ListItems retorna todos os itens armazenados no repositório
	ListItems() (MapRepo, error)
}
