FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./go.mod
COPY go.sum ./go.sum

RUN go mod download

COPY . .

RUN go build -o main .
RUN go run cmd/server/main.go

EXPOSE 8080

CMD ["./main"]