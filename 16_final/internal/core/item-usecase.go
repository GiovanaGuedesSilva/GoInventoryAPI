package core

import (
	"fmt"
	"time"

	"api/internal/core/item" // Pacote que contém a entidade Item e a interface do repositório
	// Pacote onde vivem os erros globais, como config.ErrNotFound
)

/*
ItemUsecase representa o caso de uso relacionado à entidade Item.

Ele orquestra a lógica de negócio (ex: validações, decisões, erros) e
se comunica com o repositório apenas através da interface ItemRepositoryPort.
*/
type ItemUsecase struct {
	repo item.ItemRepositoryPort // Abstração do repositório (MySQL, memória, etc.)
}

/*
NewItemUsecase cria uma nova instância do caso de uso, injetando o repositório.

Parâmetro:
- repo: uma implementação concreta de ItemRepositoryPort (ex: memória, MySQL)

Retorna:
- ItemUsecasePort (interface da aplicação)
*/
func NewItemUsecase(repo item.ItemRepositoryPort) ItemUsecasePort {
	return &ItemUsecase{
		repo: repo,
	}
}

/*
SaveItem salva um novo item, repassando a chamada para o repositório.

Pode aplicar validações de negócio antes de salvar (se necessário).

Retorna:
- Erro encadeado com contexto, caso ocorra problema no repositório.
*/
func (u *ItemUsecase) SaveItem(it item.Item) error {
	// Inicializa os timestamps
	now := time.Now()
	it.CreatedAt = now
	it.UpdatedAt = now

	if err := u.repo.SaveItem(&it); err != nil {
		return fmt.Errorf("error saving item: %w", err)
	}
	return nil
}

/*
ListItems lista todos os itens disponíveis no repositório.

Regra adicional:
- Se nenhum item for encontrado, retorna config.ErrNotFound.

Retorna:
- Mapa de itens e erro (caso ocorra)
*/
func (u *ItemUsecase) ListItems() (item.MapRepo, error) {
	its, err := u.repo.ListItems()
	if err != nil {
		return nil, fmt.Errorf("error in repository: %w", err)
	}
	if len(its) == 0 {
		// return nil, config.ErrNotFound
		// Quero que retorne a lista mesmo se estiver vazia (não é erro não ter itens)

	}
	return its, nil
}

/*
UpdateItem atualiza os dados de um item existente.

Regra de negócio pode ser aplicada aqui antes de chamar o repositório.

Retorna:
- Erro encadeado com contexto, se houver falha.
*/
func (u *ItemUsecase) UpdateItem(it item.Item) error {
	// Atualiza o timestamp de modificação
	it.UpdatedAt = time.Now()

	if err := u.repo.UpdateItem(&it); err != nil {
		return fmt.Errorf("error updating item: %w", err)
	}
	return nil
}

/*
DeleteItem remove um item com base no ID.

Regra de negócio pode ser aplicada aqui (ex: não deletar item com status X).

Retorna:
- Erro encadeado com contexto, se houver falha.
*/
func (u *ItemUsecase) DeleteItem(id int) error {
	if err := u.repo.DeleteItem(id); err != nil {
		return fmt.Errorf("error deleting item: %w", err)
	}
	return nil
}
