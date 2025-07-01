package item

import (
	"database/sql"
)

/*
mysqlRepository é uma implementação da interface ItemRepositoryPort usando MySQL como backend.

Esta estrutura encapsula a conexão com o banco de dados e implementa métodos para salvar, listar,
atualizar e deletar itens.
*/
type mysqlRepository struct {
	db *sql.DB // Conexão ativa com o banco de dados MySQL
}

/*
NewMySqlRepository retorna uma nova instância do repositório MySQL.

Parâmetros:
- db: conexão já estabelecida com o banco de dados.

Retorna:
- A interface ItemRepositoryPort (abstração), com implementação MySQL concreta.
*/
func NewMySqlRepository(db *sql.DB) ItemRepositoryPort {
	return &mysqlRepository{
		db: db,
	}
}

/*
SaveItem insere um novo item na tabela `items`.

Campos:
- code, title, description, price, stock, status, created_at, updated_at

Retorna:
- Um erro, caso a inserção falhe.
*/
func (r *mysqlRepository) SaveItem(it *Item) error {
	query := `
		INSERT INTO items 
		(code, title, description, price, stock, status, created_at, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query,
		it.Code, it.Title, it.Description,
		it.Price, it.Stock, it.Status,
		it.CreatedAt, it.UpdatedAt,
	)
	return err
}

/*
ListItems busca todos os itens da tabela `items` e os retorna em um MapRepo (mapa de itens por ID).

Retorna:
- O mapa de itens
- Um erro, caso a query falhe
*/
func (r *mysqlRepository) ListItems() (MapRepo, error) {
	query := `
		SELECT id, code, title, description, price, stock, status, created_at, updated_at 
		FROM items`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make(MapRepo)
	for rows.Next() {
		var it Item
		if err := rows.Scan(
			&it.ID, &it.Code, &it.Title, &it.Description,
			&it.Price, &it.Stock, &it.Status,
			&it.CreatedAt, &it.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items[it.ID] = it
	}

	return items, nil
}

/*
UpdateItem atualiza os dados de um item existente baseado no ID.

Campos atualizados:
- code, title, description, price, stock, status, updated_at

Retorna:
- Um erro caso o update falhe.
*/
func (r *mysqlRepository) UpdateItem(it *Item) error {
	query := `
		UPDATE items SET 
			code=?, title=?, description=?, price=?, stock=?, status=?, updated_at=?
		WHERE id=?`
	_, err := r.db.Exec(query,
		it.Code, it.Title, it.Description,
		it.Price, it.Stock, it.Status,
		it.UpdatedAt, it.ID,
	)
	return err
}

/*
DeleteItem remove um item da tabela `items` com base no ID.

Retorna:
- Um erro, caso a exclusão falhe.
*/
func (r *mysqlRepository) DeleteItem(id int) error {
	query := `DELETE FROM items WHERE id=?`
	_, err := r.db.Exec(query, id)
	return err
}
