/*
Etapa 15 - Dockerização e Inicialização Controlada com wait-for-it.sh

Objetivo:
---------
Permitir que a API seja executada de forma automática e confiável em um ambiente containerizado.
Com o uso de Docker e do script wait-for-it.sh, garantimos que a API só inicie após o banco de dados MySQL estar pronto para receber conexões.

Principais mudanças em relação à Etapa 14:
------------------------------------------
1. Adição do Dockerfile:
  - Compila a aplicação Go com `go build`.
  - Copia todos os arquivos necessários (código, módulos e script).
  - Usa Alpine Linux como base para manter a imagem leve.

2. Criação do script wait-for-it.sh:
  - Script em Bash que aguarda a disponibilidade de um host:porta antes de iniciar a aplicação.
  - Evita que a API falhe ao tentar se conectar ao banco de dados antes que ele esteja pronto.

3. Criação do docker-compose.yml:
  - Define os serviços: app (API), mysql (banco de dados), phpmyadmin (interface web).
  - Monta volumes para persistência de dados e inicialização do banco com script SQL.
  - Expõe portas 8080 (API) e 8081 (phpMyAdmin).

4. Adição do arquivo init.sql:
  - Cria o banco de dados `inventory`.
  - Cria a tabela `items`.
  - Cria o usuário `api_user` com permissões completas no banco.

5. Execução padronizada via Docker:
  - Comando único para subir toda a stack:
    > docker-compose up --build

Uso do script wait-for-it.sh:
-----------------------------
Esse script foi copiado para a raiz da aplicação e é executado no CMD do Dockerfile da seguinte forma:

CMD ["/wait-for-it.sh", "mysql:3306", "--", "/app/bin/myapp"]

Isso faz com que:
- A aplicação aguarde até que o banco de dados esteja escutando na porta 3306.
- Após a conexão ser bem-sucedida, o binário da aplicação Go seja executado.
- Em caso de falha, a aplicação não sobe de forma prematura.

Testes da API:
--------------
POST http://localhost:8080/items
Content-Type: application/json

	{
	  "id": 1,
	  "code": "ABC123",
	  "title": "Notebook",
	  "description": "Notebook Dell",
	  "price": 3500.50,
	  "stock": 10,
	  "status": "disponível"
	}

GET http://localhost:8080/items

Acesso ao banco via navegador:
------------------------------
phpMyAdmin: http://localhost:8081
Usuário: root
Senha: root

Conclusão:
----------
Essa etapa garante um ambiente de desenvolvimento robusto, padronizado e portável.
A API está agora pronta para ser usada por qualquer membro da equipe com um único comando Docker,
sem depender de instalações manuais de banco de dados, dependências Go ou configurações locais.
*/
package main
