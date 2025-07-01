-- Cria o banco de dados 'inventory' caso ainda não exista
CREATE DATABASE IF NOT EXISTS inventory;

-- Usa o banco de dados 'inventory' para as próximas instruções
USE inventory;

-- Cria a tabela 'items' com os campos necessários para o inventário
CREATE TABLE IF NOT EXISTS items (
    id INT AUTO_INCREMENT PRIMARY KEY,                         -- ID autoincrementável como chave primária
    code VARCHAR(255) NOT NULL,                                -- Código único do item
    title VARCHAR(255) NOT NULL,                               -- Título ou nome do item
    description TEXT,                                          -- Descrição longa (opcional)
    price DECIMAL(10, 2),                                      -- Preço com duas casas decimais
    stock INT,                                                 -- Quantidade em estoque
    status VARCHAR(50),                                        -- Status do item (ex: ativo, inativo, etc.)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,            -- Data de criação (valor padrão: agora)
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- Atualiza sempre que o registro é alterado
);

-- Cria o usuário 'api_user' com a senha 'api_password', acessível de qualquer host
CREATE USER 'api_user'@'%' IDENTIFIED BY 'api_password';

-- Concede todos os privilégios no banco 'inventory' ao usuário criado
GRANT ALL PRIVILEGES ON inventory.* TO 'api_user'@'%';

-- Atualiza a tabela interna de permissões do MySQL
FLUSH PRIVILEGES;
