# Etapa de construção
FROM golang:1.20-alpine AS builder

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar os arquivos go.mod e go.sum para o diretório de trabalho
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod download

# Copiar o código do projeto para o diretório de trabalho
COPY . .

# Compilar a aplicação
RUN go build -o server ./cmd/server/main.go

# Etapa final
FROM alpine:latest

# Instalar dependências necessárias para o contêiner Alpine
RUN apk --no-cache add ca-certificates

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /root/

# Copiar o binário compilado da etapa anterior
COPY --from=builder /app/server .

# Copiar o arquivo .env para o contêiner
COPY .env .

# Expor a porta da aplicação
EXPOSE 8080

# Comando para iniciar a aplicação
CMD ["./server"]
