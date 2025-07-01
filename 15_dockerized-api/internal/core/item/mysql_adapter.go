package item

import (
	"database/sql"
)

/*
mysqlRepository é a implementação do repositório de itens usando um banco de dados MySQL.

Ele implementa a interface ItemRepositoryPort, o que permite que essa implementação
seja utilizada pela camada de caso de uso de forma desacoplada.
*/
type mysqlRepository struct {
	db *sql.DB // Conexão com o banco de dados MySQL
}

/*
NewMySqlRepository cria uma nova instância do repositório MySQL.

Recebe um *sql.DB (já conectado) e retorna um objeto que implementa ItemRepositoryPort.
*/
func NewMySqlRepository(db *sql.DB) ItemRepositoryPort {
	return &mysqlRepository{
		db: db,
	}
}

/*
SaveItem insere um novo item na tabela `items` do banco de dados MySQL.

Utiliza uma query preparada com placeholders (?) para prevenir SQL Injection.
Os valores são extraídos do objeto `Item` passado como argumento.

Retorna um erro se a operação falhar.
*/
func (r *mysqlRepository) SaveItem(it *Item) error {
	query := `
		INSERT INTO items (code, title, description, price, stock, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query,
		it.Code,
		it.Title,
		it.Description,
		it.Price,
		it.Stock,
		it.Status,
		it.CreatedAt,
		it.UpdatedAt,
	)
	return err
}

/*
ListItems busca todos os itens da tabela `items` no banco MySQL.

1. Executa uma query SELECT.
2. Itera sobre os resultados usando `rows.Next()`.
3. Para cada linha, preenche uma struct `Item` usando `rows.Scan`.
4. Armazena cada item no map `MapRepo`, com a chave sendo o ID do item.

Retorna:
- Um map[int]Item com todos os registros encontrados.
- Um erro, se houver falha na consulta ou no scan.
*/
func (r *mysqlRepository) ListItems() (MapRepo, error) {
	query := `
		SELECT id, code, title, description, price, stock, status, created_at, updated_at
		FROM items
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := make(MapRepo) // Inicializa o mapa
	for rows.Next() {
		var it Item
		// Preenche a struct com os dados do banco
		if err := rows.Scan(
			&it.ID,
			&it.Code,
			&it.Title,
			&it.Description,
			&it.Price,
			&it.Stock,
			&it.Status,
			&it.CreatedAt,
			&it.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items[it.ID] = it
	}

	return items, nil
}
