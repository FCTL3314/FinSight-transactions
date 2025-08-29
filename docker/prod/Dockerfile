FROM python:3.11-slim as builder

WORKDIR /app

RUN pip install poetry

COPY poetry.lock pyproject.toml ./
RUN poetry config virtualenvs.create false && poetry install --no-dev --no-interaction --no-ansi

COPY . .

FROM python:3.11-slim

WORKDIR /app

COPY --from=builder /app /app

COPY --from=builder /go/bin/goose /usr/local/bin/
COPY ./docker/prod/scripts/entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

EXPOSE ${INTERNAL_PORT}

ENTRYPOINT ["/app/entrypoint.sh"]
