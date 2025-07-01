# internal/

O diretório `internal/` contém a lógica privada da aplicação. Em Go, qualquer pacote dentro de `internal/` não pode ser importado por outros módulos, garantindo encapsulamento e separação de responsabilidades.

Este projeto segue princípios da arquitetura limpa, com separação clara entre domínio (regras de negócio) e infraestrutura.

---

## 📁 core/

Responsável pela **regra de negócio** da aplicação.

### Subpastas e arquivos:
- `item/`: contém a entidade principal `Item`, suas portas (interfaces) e implementações de adaptadores em memória e MySQL.
- `item-usecase.go` / `item-usecase_port.go`: definição e implementação dos casos de uso relacionados ao item.

```bash
internal/core/
├── item/
│   ├── item.go              # entidade Item
│   ├── item_ports.go        # interfaces (ports)
│   ├── inmemory_adapter.go  # implementação em memória
│   └── mysql_adapter.go     # implementação com MySQL
├── item-usecase.go          # caso de uso principal
├── item-usecase_port.go     # interface do caso de uso
```

---

## 📁 platform/

Contém a **infraestrutura** da aplicação – implementações específicas de acesso a dados.

- `mysql/`: configuração do MySQL
- `mongodb/`: configuração do MongoDB

```bash
internal/platform/
├── mysql/
│   └── mysql-setup.go       # setup de conexão com MySQL
└── mongodb/
    └── mongodb.go           # setup de conexão com MongoDB
```

---

## ✅ Benefícios dessa estrutura

- **Separação de responsabilidades:** domínio isolado da infraestrutura.
- **Escalabilidade:** fácil adicionar novos bancos ou entidades.
- **Testabilidade:** adaptadores podem ser trocados por mocks ou fakes.
- **Modularidade:** fácil manter e entender cada parte.

---