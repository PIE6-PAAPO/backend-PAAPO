# Etapa 1: Build
FROM golang:1.22-alpine AS builder

ENV GOTOOLCHAIN=auto

WORKDIR /app

# Copiar arquivos de dependência primeiro (cache de camadas)
COPY go.mod go.sum ./
RUN go mod download

# Copiar o restante da aplicação
COPY . .

# Compilar o binário
RUN go build -o main ./cmd/server

# Etapa 2: Runtime
FROM alpine:3.20

WORKDIR /app

# Copiar apenas o binário da etapa anterior
COPY --from=builder /app/main .

# Expor a porta do servidor
EXPOSE 8080

# Comando padrão
CMD ["./main"]