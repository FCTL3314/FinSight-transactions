FROM ghcr.io/astral-sh/uv:python3.13-trixie-slim

RUN apt-get update && apt-get install -y --no-install-recommends curl make

ENV PATH="/root/.local/bin:$PATH"

COPY . /app
WORKDIR /app

RUN uv sync --locked

EXPOSE ${INTERNAL_PORT}

ENTRYPOINT ["/app/docker/local/scripts/entrypoint.sh"]
