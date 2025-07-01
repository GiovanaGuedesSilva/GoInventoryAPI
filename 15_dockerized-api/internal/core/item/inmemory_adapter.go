package item

import (
	"fmt"
)

/*
MapRepository é uma implementação **em memória** do repositório de itens.
Ele usa um `map` para armazenar os itens em tempo de execução, sem necessidade de banco de dados.
É muito útil para testes locais ou quando queremos rodar a aplicação sem infraestrutura externa.
*/
type MapRepository struct {
	items MapRepo // Mapa onde os itens são armazenados. MapRepo é um alias para map[int]Item
}

/*
NewMapRepository cria uma nova instância do repositório em memória.

Retorna: um ponteiro para MapRepository como `ItemRepositoryPort` (interface).
Isso permite que ele seja usado no lugar do repositório MySQL, graças ao polimorfismo.
*/
func NewMapRepository() ItemRepositoryPort {
	return &MapRepository{
		items: make(MapRepo), // Inicializa o mapa vazio
	}
}

/*
SaveItem salva um novo item no repositório em memória.

Regras:
- O ID não pode ser 0.
- Não pode haver dois itens com o mesmo ID.

Se as validações forem passadas, o item é salvo no mapa.
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
ListItems retorna todos os itens armazenados no repositório em memória.

Retorna:
- Um MapRepo (map[int]Item)
- Nil como erro, pois essa implementação não possui falhas esperadas
*/
func (r *MapRepository) ListItems() (MapRepo, error) {
	return r.items, nil
}
