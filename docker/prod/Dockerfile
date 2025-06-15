FROM golang:1.24-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download & \
    go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o transactions_app ./cmd/api

FROM alpine:3.22

COPY --from=builder /app/transactions_app /usr/local/bin/transactions_app

CMD ["transactions_app"]

EXPOSE 8080