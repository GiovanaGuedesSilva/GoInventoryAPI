package gosqldriver

import (
	"fmt"
)

/*
MySQLClientConfig contém os dados necessários para se conectar a uma base de dados MySQL.

Essa configuração é usada para montar a DSN (Data Source Name), que é a string
utilizada pelo driver `database/sql` para se conectar ao MySQL.
*/
type MySQLClientConfig struct {
	User     string // Nome do usuário do banco de dados
	Password string // Senha do usuário
	Host     string // Endereço do host onde o MySQL está rodando (ex: localhost, mysql)
	Port     string // Porta de conexão (normalmente 3306)
	Database string // Nome do banco de dados a ser utilizado
}

/*
dsn gera a string de conexão (Data Source Name) no formato esperado pelo driver do MySQL.

Formato gerado:

	usuario:senha@tcp(host:porta)/banco?charset=utf8mb4&parseTime=True&loc=Local

- `utf8mb4`: permite suporte a emojis e caracteres especiais
- `parseTime=True`: faz com que o Go trate campos de data/hora corretamente
- `loc=Local`: define o fuso horário como local (pode ser ajustado conforme necessário)

Essa função é usada internamente no método `connect()` do `MySQLClient`.
*/
func (config MySQLClientConfig) dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.Host, config.Port, config.Database)
}
