package core

import (
	"fmt"

	"api/internal/core/item"
	"api/pkg/config"
)

/*
ItemUsecase representa a implementação concreta da interface ItemUsecasePort.

É nesta camada que colocamos as regras de negócio da aplicação.
Ela orquestra as chamadas ao repositório, validando ou processando dados conforme necessário.
*/
type ItemUsecase struct {
	repo item.ItemRepositoryPort // Repositório de itens (MySQL, memória, etc.)
}

/*
NewItemUsecase cria uma nova instância do caso de uso de item.

Recebe como dependência um repositório que implementa ItemRepositoryPort.
Isso permite flexibilidade (injeção de dependência) e facilita testes.
*/
func NewItemUsecase(repo item.ItemRepositoryPort) ItemUsecasePort {
	return &ItemUsecase{
		repo: repo,
	}
}

/*
SaveItem salva um novo item utilizando o repositório.

Regras atuais:
- Encapsula e propaga qualquer erro retornado pelo repositório.
- Usa `%w` no `fmt.Errorf` para manter a rastreabilidade do erro original.

Pode ser expandido para:
- Validações de campos obrigatórios
- Regras de negócio (ex: não permitir estoque negativo)
*/
func (u *ItemUsecase) SaveItem(it item.Item) error {
	if err := u.repo.SaveItem(&it); err != nil {
		return fmt.Errorf("erro ao salvar item: %w", err)
	}
	return nil
}

/*
ListItems busca todos os itens armazenados no repositório.

Regras:
- Se ocorrer erro na consulta ao repositório, ele é encapsulado com contexto.
- Se o repositório retornar um mapa vazio (nenhum item), retorna um erro padrão `config.ErrNotFound`.

Isso permite que a camada superior (ex: handler HTTP) saiba que a lista está vazia.
*/
func (u *ItemUsecase) ListItems() (item.MapRepo, error) {
	its, err := u.repo.ListItems()
	if err != nil {
		return nil, fmt.Errorf("erro no repositório: %w", err)
	}
	if len(its) == 0 {
		return nil, config.ErrNotFound
	}
	return its, nil
}
