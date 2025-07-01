/*
Etapa 13 – Separação em Arquitetura em Camadas

Objetivo:
---------
Transformar a estrutura monolítica em uma arquitetura em camadas, promovendo separação de responsabilidades e melhor organização do código, com base na arquitetura hexagonal.

Principais mudanças em relação à etapa 12:
------------------------------------------
- Criação da estrutura de diretórios:
  - cmd/: pontos de entrada para REST, CLI e gRPC;
  - internal/adapters/handlers: implementação dos handlers para diferentes interfaces;
  - internal/adapters/repositories: múltiplas implementações de persistência (memória, MongoDB, MySQL);
  - internal/domain: definição dos modelos de domínio e interfaces (contratos);
  - internal/usecases: lógica de negócio separada por entidade;
  - internal/platform: drivers específicos de infraestrutura para os bancos;

- Separação real entre lógica de aplicação, domínio e infraestrutura;
- Introdução de suporte para múltiplas interfaces de acesso à aplicação.

Benefícios:
-----------
- Organização por responsabilidade;
- Melhora a testabilidade e a manutenibilidade;
- Permite reutilização da lógica em diferentes entradas;
- Segue princípios da arquitetura limpa (dependências apontam para o centro).
*/
package main
