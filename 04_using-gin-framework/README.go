// readme.go
//
// Este arquivo é apenas explicativo. Ele descreve o papel dos arquivos go.mod e go.sum
// em projetos escritos em Go usando o sistema oficial de módulos (Go Modules).
//
// Autor: Giovana Guedes
// Projeto: GoInventoryAPI
// --------------------------------------------

/*
	📦 go.mod

	O go.mod é o arquivo principal de definição de módulo no Go.

	Ele serve para:
	- Declarar que o projeto é um módulo Go
	- Informar o nome do módulo (geralmente o caminho do repositório)
	- Definir a versão mínima do Go necessária
	- Listar dependências externas e suas versões

	Exemplo típico de go.mod:

		module github.com/GiovanaGuedesSilva/GoInventoryAPI

		go 1.22

		require (
			github.com/gin-gonic/gin v1.9.1
		)

	Você cria esse arquivo com:

		go mod init github.com/seu-usuario/seu-projeto

	E adiciona dependências com:

		go get github.com/algumapkg

	Ou organiza com:

		go mod tidy

	⚠️ Nunca edite esse arquivo manualmente sem saber o que está fazendo.

*/

/*
	🔐 go.sum

	O go.sum é gerado automaticamente pelo Go para garantir a integridade
	e segurança das dependências declaradas no go.mod.

	Ele armazena hashes criptográficos (checksums) que o Go usa para verificar
	se os arquivos baixados são legítimos e não foram alterados.

	⚠️ Esse arquivo:
	- **Deve** ser versionado junto com seu código (suba no Git)
	- **Não** deve ser editado manualmente
	- É usado em builds, CI/CD e ambientes de produção para garantir reprodutibilidade

	Exemplo de entrada em go.sum:

		github.com/gin-gonic/gin v1.9.1 h1:8Y92s...
		github.com/gin-gonic/gin v1.9.1/go.mod h1:kd7h...

	Se estiver com dependências bagunçadas, use:

		go mod tidy
*/

/*
	🛠️ Resumo dos comandos úteis:

	- go mod init <modulo>       → Cria o go.mod
	- go get <pacote>            → Adiciona dependência
	- go mod tidy                → Limpa e organiza as dependências
	- go list -m all             → Lista todos os módulos
	- go build                   → Compila o projeto
	- go run main.go             → Executa a aplicação

	💡 Dica:
	Mantenha sempre seu go.mod e go.sum atualizados e versionados corretamente.
	Isto é parte essencial de qualquer projeto profissional em Go.
*/

// Este arquivo não contém código executável.
// Ele existe apenas para servir como documentação embutida no projeto.
package main
