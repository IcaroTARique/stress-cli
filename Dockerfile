# Use uma imagem base oficial do Go para a construção
FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o /app/stress-cli ./main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/stress-cli .

ENTRYPOINT ["./stress-cli"]
