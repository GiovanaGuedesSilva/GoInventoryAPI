# Utilizamos a imagem oficial do Golang com Alpine (leve e otimizada para produção)
FROM golang:1.22.3-alpine3.20

# Instalamos o bash, necessário para rodar o script wait-for-it.sh
RUN apk add --no-cache bash

# Define o diretório de trabalho dentro do contêiner (tudo será executado a partir daqui)
WORKDIR /app

# Copiamos os arquivos de dependência Go (go.mod e go.sum)
COPY go.mod go.sum ./

# Faz o download antecipado dos módulos (otimiza cache em builds futuros)
RUN go mod download

# Copiamos o restante dos arquivos da aplicação para dentro do contêiner
COPY . .

# Copiamos o script wait-for-it.sh (usado para aguardar o MySQL)
COPY wait-for-it.sh /wait-for-it.sh

# Damos permissão de execução ao script e construímos o binário da aplicação
RUN chmod +x /wait-for-it.sh && \
    go build -o /app/bin/myapp ./cmd/rest/main.go

# Expomos a porta 8080 (a que a aplicação usa)
EXPOSE 8080

# Comando padrão do contêiner:
# Primeiro espera o serviço MySQL responder na porta 3306, depois executa a aplicação
CMD ["/wait-for-it.sh", "mysql:3306", "--", "/app/bin/myapp"]
