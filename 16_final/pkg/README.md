# pkg/

O diretório `pkg/` contém **pacotes reutilizáveis** que não fazem parte da lógica de negócio central, mas são usados por diferentes partes da aplicação. Esses pacotes são projetados para serem genéricos e, em teoria, poderiam ser reutilizados em outros projetos Go.

---

## 📁 config/

Contém funcionalidades relacionadas à **configuração da aplicação**, como leitura de variáveis de ambiente.

### Arquivo:
- `config.go`: carrega configurações gerais da aplicação, como porta do servidor, conexões externas etc.

---

## 📁 mysql/

Pacotes relacionados à conexão com o banco de dados MySQL.

### 📁 go-sql-driver/

Implementa a configuração e inicialização do cliente MySQL utilizando `database/sql` com o driver `go-sql-driver/mysql`.

#### Arquivos:
- `mysql-client.go`: função de inicialização do cliente MySQL.
- `mysql-config.go`: estrutura de configuração do MySQL (usuário, senha, host, etc.).

---

## Benefícios desta estrutura

- **Reutilização:** pacotes aqui podem ser usados por diferentes módulos da aplicação.
- **Organização:** separa utilitários e componentes de infraestrutura da lógica de negócio.
- **Testabilidade:** permite testar configurações de forma isolada.

---

## Estrutura visual

```bash
pkg/
├── config/
│   └── config.go
└── mysql/
    └── go-sql-driver/
        ├── mysql-client.go
        └── mysql-config.go
```