FROM golang:1.24.3-alpine

RUN apk add --no-cache git make curl
RUN go install github.com/air-verse/air@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY docker/local/scripts/entrypoint.sh /usr/local/bin/entrypoint.sh
RUN chmod +x /usr/local/bin/entrypoint.sh

EXPOSE ${INTERNAL_PORT}

ENTRYPOINT ["entrypoint.sh"]