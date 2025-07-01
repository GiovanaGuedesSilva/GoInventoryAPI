/*
Etapa 14 – Refatoração para Arquitetura Limpa

Objetivo:
---------
Aprimorar a separação de responsabilidades e aplicar de forma mais fiel os princípios da Clean Architecture.

Principais mudanças em relação à etapa 13:
------------------------------------------
- Refatoração da pasta internal:
  - internal/core/ passa a centralizar o domínio e os usecases;
  - Cada entidade possui seu próprio subdiretório (ex: core/item/) com:
  - Modelo de domínio (item.go)
  - Portas (interfaces): item_ports.go
  - Adaptadores específicos: inmemory_adapter.go, mysql_adapter.go
  - Separação entre interface (`*_port.go`) e implementação (`*_usecase.go`);

- Handlers REST movidos para cmd/rest/handlers;
- Criação da pasta pkg/config para centralização de configurações reutilizáveis;
- Maior clareza entre camadas e contratos, respeitando o princípio de dependência.

Benefícios:
-----------
- Estrutura mais alinhada à Clean Architecture;
- Facilita colaboração entre diferentes times;
- Componentes fortemente desacoplados;
- Código preparado para crescer em complexidade mantendo organização.
*/
package main
