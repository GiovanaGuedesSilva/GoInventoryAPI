package item

import (
	"fmt"
)

/*
MapRepository é uma implementação do repositório de itens em memória.

- Utiliza um `map[int]Item` para armazenar os itens durante a execução do programa.
- É útil para testes locais ou execução sem banco de dados.
*/
type MapRepository struct {
	items MapRepo // MapRepo é um alias para map[int]Item
}

/*
NewMapRepository cria uma nova instância do repositório em memória.

Retorna:
- Um ponteiro para MapRepository, mas como o tipo ItemRepositoryPort (interface).
Isso permite substituição transparente de implementações concretas.
*/
func NewMapRepository() ItemRepositoryPort {
	return &MapRepository{
		items: make(MapRepo), // Inicializa mapa vazio
	}
}

/*
SaveItem salva um novo item no repositório.

Regras:
- O ID não pode ser zero.
- Não pode existir outro item com o mesmo ID.

Retorna erro caso a operação viole alguma dessas regras.
*/
func (r *MapRepository) SaveItem(it *Item) error {
	if it.ID == 0 {
		return fmt.Errorf("ID do item não pode ser 0")
	}
	if _, exists := r.items[it.ID]; exists {
		return fmt.Errorf("já existe um item com o ID %d", it.ID)
	}
	r.items[it.ID] = *it
	return nil
}

/*
ListItems retorna todos os itens armazenados.

Retorna:
- O mapa de itens.
- `nil` como erro, pois essa operação não deve falhar nesta implementação.
*/
func (r *MapRepository) ListItems() (MapRepo, error) {
	return r.items, nil
}

/*
UpdateItem atualiza um item existente no repositório.

Regras:
- O ID não pode ser zero.
- O item deve já existir no mapa.

Retorna erro caso as validações falhem.
*/
func (r *MapRepository) UpdateItem(it *Item) error {
	if it.ID == 0 {
		return fmt.Errorf("ID do item não pode ser 0")
	}
	if _, exists := r.items[it.ID]; !exists {
		return fmt.Errorf("item com ID %d não existe", it.ID)
	}
	r.items[it.ID] = *it
	return nil
}

/*
DeleteItem remove um item do repositório em memória.

Regras:
- O ID não pode ser zero.
- O item deve existir.

Retorna erro caso as validações falhem.
*/
func (r *MapRepository) DeleteItem(id int) error {
	if id == 0 {
		return fmt.Errorf("ID do item não pode ser 0")
	}
	if _, exists := r.items[id]; !exists {
		return fmt.Errorf("item com ID %d não existe", id)
	}
	delete(r.items, id)
	return nil
}
