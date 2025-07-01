# GoInventoryAPI

API RESTful escrita em Go para gerenciamento de inventário.  
Este projeto tem como objetivo controlar produtos, quantidades e operações básicas de entrada e saída de estoque, evoluindo passo a passo com boas práticas.

---

## 🚀 Tecnologias

- [Go (Golang)](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [net/http](https://pkg.go.dev/net/http) (usado nas etapas iniciais)
- Padrão REST
- (Opcional: Docker, Banco de Dados, Swagger — em etapas futuras)

---

## 📁 Estrutura do Projeto

Cada diretório representa uma etapa evolutiva da API, com foco em modularização, boas práticas e Clean Architecture.

```bash
GoInventoryAPI/
├── step_01_basic_http/              # Servidor HTTP básico com uma rota
├── step_02_multi_endpoints/         # Múltiplas rotas com net/http
├── step_03_http_methods/            # Uso de métodos HTTP diferentes (GET, POST, etc.)
├── step_04_using_gin/               # Substituição de net/http por Gin
├── step_05_gin_with_comments/       # Mesmo código anterior, agora totalmente comentado
├── step_06_gin_improved_post/       # POST com mensagem fixa no body
├── step_07_gin_body_parsing/        # POST lendo o corpo da requisição
├── step_08_gin_json_parsing/        # POST com JSON e logging
├── step_09_handler_struct/          # Criação de struct handler com métodos associados
├── step_10_usecase_layer/           # Introdução da camada de Usecase
├── step_11_repository_layer/        # Implementação de um repositório em memória
├── step_12_interface_usecase/       # Abstração da camada de usecase via interface
├── go.mod
├── go.sum
└── README.md                        # Este arquivo
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


