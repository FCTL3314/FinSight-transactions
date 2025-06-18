FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download & \
    go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o transactions_app ./cmd/api

FROM alpine:3.22
WORKDIR /app

# TODO: Удалить строку ниже, так как .env и config.yml файлы будут создаваться уже на проде.
COPY --from=builder /app/.env /app/config.yml /app/transactions_app /app/

CMD ["/app/transactions_app"]

EXPOSE 8080