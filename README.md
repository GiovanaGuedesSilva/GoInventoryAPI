# GoInventoryAPI

API RESTful escrita em Go para gerenciamento de inventário.  
Este projeto tem como objetivo controlar produtos, quantidades e operações básicas de entrada e saída de estoque.

---

## 🚀 Tecnologias

- [Go (Golang)](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [net/http](https://pkg.go.dev/net/http) (nas versões iniciais)
- Padrão REST
- (Opcional: Docker, Banco de Dados, Swagger — se aplicável no futuro)

---

## 📁 Estrutura do Projeto

```bash
GoInventoryAPI/
├── step_01_basic_http/          # Servidor HTTP básico com uma rota
├── step_02_multi_endpoints/     # Várias rotas usando net/http
├── step_03_http_methods/        # Suporte a múltiplos métodos HTTP
├── step_04_using_gin/           # Introdução ao framework Gin
├── README.md                    # Este arquivo
└── ...                          # Próximas etapas evoluindo a API
