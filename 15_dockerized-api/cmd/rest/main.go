package main

import (
	"log"

	"github.com/gin-gonic/gin"

	handler "api/cmd/rest/handlers"          // Pacote responsável por lidar com requisições HTTP
	core "api/internal/core"                 // Camada de lógica de negócio
	item "api/internal/core/item"            // Pacote com o modelo e repositórios de Item
	mysqlsetup "api/internal/platform/mysql" // Configuração do cliente MySQL
)

func main() {
	/*
		Configura a conexão com o banco de dados MySQL.
		A função NewMySQLSetup encapsula toda a lógica de conexão.
		Se der erro, o programa encerra com uma mensagem.
	*/
	mysqlClient, err := mysqlsetup.NewMySQLSetup()
	if err != nil {
		log.Fatalf("Não foi possível configurar o MySQL: %v", err)
	}
	// Fecha a conexão com o banco ao encerrar o programa
	defer mysqlClient.Close()

	/*
		Inicializa o repositório de itens usando o banco MySQL.
		Esse repositório implementa a interface necessária para a camada de caso de uso.
	*/
	repo := item.NewMySqlRepository(mysqlClient.DB())

	/*
		Alternativamente, poderíamos usar um repositório em memória (útil para testes locais),
		sem necessidade de banco de dados.

		Exemplo:
		repo := item.NewMapRepository()
	*/

	/*
		Cria o caso de uso da aplicação, que contém a lógica de negócio.
		Recebe o repositório como dependência (injeção de dependência).
	*/
	usecase := core.NewItemUsecase(repo)

	/*
		Cria o handler responsável por expor os endpoints HTTP,
		se comunicando com a lógica de negócio via o caso de uso.
	*/
	handler := handler.NewHandler(usecase)

	/*
		Configura o roteador HTTP usando o framework Gin.
		Define as rotas disponíveis para o cliente.
	*/
	router := gin.Default()
	router.POST("/items", handler.SaveItem) // Rota para salvar um item
	router.GET("/items", handler.ListItems) // Rota para listar todos os itens

	// Inicia o servidor web na porta 8080
	log.Println("Servidor iniciado em http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
