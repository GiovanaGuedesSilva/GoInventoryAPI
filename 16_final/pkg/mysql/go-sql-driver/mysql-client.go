package gosqldriver

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
MySQLClient representa um cliente para interagir com uma base de dados MySQL.

Ele encapsula:
- A configuração da conexão (host, usuário, senha, etc.)
- A instância da conexão ativa com o banco (`*sql.DB`)
*/
type MySQLClient struct {
	config MySQLClientConfig // Configuração do cliente MySQL (ver struct abaixo)
	db     *sql.DB           // Conexão ativa com o banco de dados
}

/*
NewMySQLClient cria uma nova instância de MySQLClient e estabelece a conexão com o banco.

Fluxo:
1. Cria o struct com base na configuração recebida.
2. Chama `connect()` para abrir e testar a conexão com o banco.
3. Se bem-sucedido, retorna o cliente; caso contrário, retorna erro.

Essa função centraliza a criação segura do cliente MySQL.
*/
func NewMySQLClient(config MySQLClientConfig) (*MySQLClient, error) {
	client := &MySQLClient{config: config}
	err := client.connect()
	if err != nil {
		return nil, fmt.Errorf("falha ao inicializar o MySQLClient: %v", err)
	}
	return client, nil
}

/*
connect realiza a conexão com o banco de dados MySQL com base na configuração.

Passos:
- Monta a DSN (data source name) com `config.dsn()`.
- Usa `sql.Open` para abrir a conexão.
- Realiza um `Ping()` para garantir que a conexão está funcional.

Se der erro em qualquer etapa, ele é retornado com contexto.
*/
func (client *MySQLClient) connect() error {
	dsn := client.config.dsn() // Monta string de conexão
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("falha ao conectar ao MySQL: %w", err)
	}
	if err := conn.Ping(); err != nil {
		return fmt.Errorf("falha ao verificar conexão com MySQL: %w", err)
	}
	client.db = conn
	return nil
}

/*
Close fecha a conexão com o banco de dados, se estiver aberta.

É importante sempre chamar essa função no final do ciclo de vida da aplicação
(geralmente com `defer`) para liberar os recursos.
*/
func (client *MySQLClient) Close() {
	if client.db != nil {
		client.db.Close()
	}
}

/*
DB retorna o ponteiro para `*sql.DB`, permitindo que outras partes da aplicação
(usecases, repositórios) executem queries SQL com segurança.
*/
func (client *MySQLClient) DB() *sql.DB {
	return client.db
}
