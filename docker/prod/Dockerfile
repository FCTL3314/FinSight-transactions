FROM golang:1.24.3-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o app ./cmd/api

FROM alpine:latest

RUN apk add --no-cache bash ca-certificates make git curl go \
    && go install github.com/pressly/goose/v3/cmd/goose@latest

RUN adduser -D -g '' appuser

WORKDIR /app

COPY --from=builder /app/app .
COPY .env /app/.env
COPY config.yml /app/config.yml
COPY docker/prod/scripts/entrypoint.sh /app/entrypoint.sh
COPY Makefile /app/Makefile

RUN chmod +x /app/entrypoint.sh

USER appuser

EXPOSE 8080

ENTRYPOINT ["/app/entrypoint.sh"]
