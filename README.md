# GoInventoryAPI

API RESTful escrita em Go para gerenciamento de inventário.  
Este projeto tem como objetivo controlar produtos, quantidades e operações básicas de entrada e saída de estoque, evoluindo passo a passo com boas práticas.

---

## 🚀 Tecnologias

- [Go (Golang)](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [net/http](https://pkg.go.dev/net/http) (usado nas etapas iniciais)
- Banco de dados MySQL
- Docker e Docker Compose
- phpMyAdmin (interface web para o MySQL)

---

## Pré-requisitos

1. Docker instalado em seu sistema.  
2. Docker Compose instalado em seu sistema.  
3. Conexão com a internet para baixar as imagens necessárias do Docker.

## Conteúdo do Repositório

- `Dockerfile`: Arquivo de configuração para construir a imagem da aplicação em Golang.  
- `docker-compose.yml`: Arquivo de configuração para orquestrar os serviços Docker (aplicação, MySQL e phpMyAdmin).  
- `init.sql`: Script SQL para inicializar o banco de dados MySQL com o esquema necessário e o usuário da API.  
- Código-fonte da API de Inventário.


## Instruções de Configuração

### Passo 1: Configurar o Banco de Dados

Antes de iniciar os serviços Docker, é necessário garantir que o script `init.sql` seja executado para configurar o banco de dados.  
Esse script cria o banco de dados `inventory`, a tabela `items` e um usuário da API com as permissões adequadas.

### Passo 2: Iniciar os Serviços com Docker

Use o Docker Compose para iniciar todos os serviços definidos no arquivo `docker-compose.yml`:

```sh
docker-compose up --build
```

Esse comando fará o seguinte:

1. Construirá a imagem da aplicação em Golang.  
2. Iniciará o contêiner do MySQL.  
3. Iniciará o contêiner do phpMyAdmin.  
4. Iniciará o contêiner da aplicação Golang.

### Passo 3: Executar o Script SQL no phpMyAdmin

Abra seu navegador e acesse [http://localhost:8081](http://localhost:8081) para abrir o phpMyAdmin.  
Use as seguintes credenciais para login:

- **Usuário:** `root`  
- **Senha:** `root`

Depois de acessar:

1. Selecione o banco de dados `inventory`.  
2. Vá até a aba "SQL".  
3. Copie e cole o conteúdo do arquivo `init.sql`.  
4. Execute o script.

### Passo 4: Verificar o Funcionamento da API

Com todos os contêineres em execução e o banco de dados configurado, acesse [http://localhost:8080](http://localhost:8080) no navegador, ou utilize ferramentas como `curl` ou `Postman` para interagir com os endpoints `/items`.

## Endpoints da API

### `POST /items` - Criar um novo item no inventário

Exemplo de corpo JSON:

```json
{
  "id": 1,
  "code": "ITEM001",
  "title": "Example Item",
  "description": "This is an example item",
  "price": 29.99,
  "stock": 50,
  "status": "available",
  "created_at": "2024-07-17T15:04:05Z",
  "updated_at": "2024-07-17T15:04:05Z"
}
```

### `GET /items` - Obter todos os itens do inventário

## Exemplos de Uso com `curl`

### Criar um novo item

```sh
curl -X POST http://localhost:8080/items -H "Content-Type: application/json" -d '{
  "id": 1,
  "code": "ITEM001",
  "title": "Example Item",
  "description": "This is an example item",
  "price": 29.99,
  "stock": 50,
  "status": "available",
  "created_at": "2024-07-17T15:04:05Z",
  "updated_at": "2024-07-17T15:04:05Z"
}'
```

### Obter a lista de itens

```sh
curl http://localhost:8080/items
```

## Solução de Problemas

### Erro de conexão com o MySQL

Verifique se:

1. O contêiner do MySQL está em execução.  
2. O script `init.sql` foi executado corretamente.  
3. As credenciais de banco de dados no código coincidem com as do script SQL.

### Verificar logs dos contêineres

```sh
docker-compose logs app
docker-compose logs mysql
docker-compose logs phpmyadmin
```

## Conclusão

Seguindo esses passos, você conseguirá configurar e executar corretamente a API de Inventário.  
Se encontrar problemas, consulte os logs dos contêineres e verifique se todos os serviços estão configurados corretamente.

## 📁 Estrutura do Projeto

Cada diretório representa uma etapa evolutiva da API, com foco em modularização, boas práticas e Clean Architecture.

```bash
C:\GoProjects\GoInventoryAPI
│
├── .github/                              # Configurações e workflows do GitHub Actions
├── 01_basic-http-server/                 # Servidor HTTP básico com uma rota fixa
├── 02_multi_endpoints/                   # Múltiplos endpoints usando net/http
├── 03_http_methods/                      # Uso de diferentes métodos HTTP (GET, POST, etc.)
├── 04_using-gin-framework/               # Substituição do net/http pelo framework Gin
├── 05_adding-comments/                   # Código da etapa anterior com comentários explicativos
├── 06_handler_logic/                     # Introdução de lógica de manipulação no handler
├── 07_read_body_plain/                   # Leitura do corpo da requisição como texto plano
├── 08_json_input/                        # Leitura e parse de JSON via body da requisição
├── 09_struct_handlers/                   # Refatoração com uso de struct handler e métodos
├── 10_usecase_domain/                    # Introdução da camada de domínio (usecase)
├── 11_query-param-handling/              # Manipulação de parâmetros de query na URL
├── 12_usecase_abstraction/               # Abstração da camada de usecase com interface
├── 13_layered-architecture-separation/   # Separação em camadas: handler, usecase, repository
├── 14_clean-architecture-refactor/       # Refatoração para aderir à Clean Architecture
├── 15_dockerized-api/                    # Containerização da API com Docker e Docker Compose
├── 16_final/                             # Versão final consolidada da API
├── go.mod                                # Arquivo de definição de módulos Go
├── go.sum                                # Checksum das dependências Go
└── README.md                             # Documentação principal do projeto

```

---

## ⚙️ Como rodar cada etapa

1. **Acesse o diretório da etapa desejada:**
   ```bash
   cd step_04_using_gin
   ```

2. **Execute a aplicação:**
   ```bash
   go run main.go
   ```

3. **Se for a primeira execução e faltar alguma dependência (como Gin):**
   ```bash
   go get github.com/gin-gonic/gin
   ```

---

## 📦 Sobre os arquivos `go.mod` e `go.sum`

O GoInventoryAPI já inclui os arquivos `go.mod` e `go.sum` em cada etapa para garantir que as dependências sejam resolvidas corretamente.

Esses arquivos são gerados automaticamente pelo Go quando você trabalha com módulos, e servem para gerenciar as dependências do projeto.

### 📄 `go.mod`

- Define o nome do módulo (nome do projeto).
- Lista os pacotes externos que o código importa diretamente.
- É gerado com o comando:
  ```bash
  go mod init <nome-do-modulo>
  ```

### 📄 `go.sum`

- Contém **hashes criptográficos** para garantir a integridade e segurança das dependências.
- É criado automaticamente quando você executa:
  ```bash
  go build
  go run
  go get
  ```
- Não deve ser editado manualmente — o Go gerencia esse arquivo por você.

> ✅ Como esses arquivos já estão versionados no repositório, **não é necessário gerar nada manualmente**. Basta rodar `go run main.go` na etapa desejada e o projeto funcionará normalmente.


