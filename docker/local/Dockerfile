FROM golang:1.24.3-alpine

RUN apk add --no-cache git make curl

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY docker/local/scripts/entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

ENTRYPOINT ["/app/entrypoint.sh"]