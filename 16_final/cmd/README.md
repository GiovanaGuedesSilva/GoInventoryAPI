# cmd/

O diretório `cmd/` contém os pontos de entrada da aplicação. Cada subdiretório representa uma **variante executável** da aplicação Go. Essa é uma convenção adotada em projetos maiores e profissionais para manter os *entrypoints* bem organizados.

A separação entre os modos de execução permite que diferentes formas de rodar a aplicação (como uma API REST, um serviço gRPC ou uma ferramenta de linha de comando) tenham inicialização própria, sem misturar responsabilidades.

## Subdiretórios

### 📁 cli/

Este subdiretório é responsável por um executável de **linha de comando** (Command Line Interface).

- **Uso comum:** scripts administrativos, tarefas de manutenção, importação/exportação de dados, verificação de status etc.
- **Tecnologias típicas:** pode usar bibliotecas como [`cobra`](https://github.com/spf13/cobra) ou `urfave/cli`.

**Exemplo:**  
```bash
go run cmd/cli/main.go --help
```

### 📁 rest/

Este diretório contém a versão API RESTful da aplicação, usando HTTP como protocolo de comunicação.

- **Uso:** é onde você inicializa seu framework web (como Gin, Echo, Fiber, etc.).
-**Responsabilidades típicas:** carregar configurações, montar rotas, injetar dependências e iniciar o servidor HTTP.

**Exemplo:**  
```bash
go run cmd/rest/main.go
```

### 📁 grpc/

Este diretório serve como ponto de entrada para o servidor gRPC.

- **Uso:** serviços que usam o protocolo gRPC (baseado em HTTP/2), geralmente para comunicação entre serviços (microservices).
-**Responsabilidades:** carregar protos compilados, inicializar o servidor gRPC, registrar serviços e iniciar o listener.

**Exemplo:**  
```bash
go run cmd/grpc/main.go
```

## Benefícios desta organização

- **Escalabilidade:** você pode manter diferentes tipos de execução sem misturar responsabilidades.
- **Isolamento:** cada modo (REST, CLI, gRPC) tem suas dependências e inicializações.
- **Clareza:** ao abrir o projeto, fica fácil identificar os pontos de entrada e os modos de uso da aplicação.
- **Modularidade:** permite testar, buildar e implantar partes específicas separadamente.