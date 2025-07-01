package mysqlsetup

import (
	gosqldriver "api/pkg/mysql/go-sql-driver"
)

/*
NewMySQLSetup configura e retorna um novo cliente MySQL.

Essa função encapsula a criação do cliente de banco de dados.
É usada pelo `main.go` para obter uma conexão pronta para uso.

Retorno:
- Um ponteiro para `MySQLClient` (estrutura que provavelmente encapsula `*sql.DB`)
- Um erro, caso a conexão falhe
*/
func NewMySQLSetup() (*gosqldriver.MySQLClient, error) {
	// Define as credenciais e dados de conexão com o banco MySQL
	config := gosqldriver.MySQLClientConfig{
		User:     "api_user",     // Nome do usuário do banco
		Password: "api_password", // Senha do banco
		Host:     "mysql",        // Host (nome do serviço Docker ou IP)
		Port:     "3306",         // Porta padrão do MySQL
		Database: "inventory",    // Nome do banco de dados a ser usado
	}

	// Cria o cliente usando o pacote go-sql-driver
	return gosqldriver.NewMySQLClient(config)
}
