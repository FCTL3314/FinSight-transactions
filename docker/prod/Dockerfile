FROM ghcr.io/astral-sh/uv:python3.13-trixie-slim AS builder

ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

WORKDIR /app

COPY pyproject.toml uv.lock ./
RUN uv sync --locked --no-dev

COPY . .

FROM ghcr.io/astral-sh/uv:python3.13-trixie-slim

ENV PATH="/root/.local/bin:$PATH"
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1
ENV VENV_PATH=/opt/venv
ENV PATH="$VENV_PATH/bin:$PATH"
ENV APP_USER=appuser

RUN apt-get update && \
    apt-get install -y make && \
    rm -rf /var/lib/apt/lists/*

RUN groupadd -r ${APP_USER} && \
    useradd --no-log-init -r -g ${APP_USER} ${APP_USER}

WORKDIR /app

COPY --from=builder ${VENV_PATH} ${VENV_PATH}
COPY --from=builder /app /app

ENV UV_CACHE_DIR=/tmp/uv_cache
COPY ./docker/prod/scripts/entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh
RUN chown -R ${APP_USER}:${APP_USER} /app

USER ${APP_USER}

EXPOSE ${INTERNAL_PORT}

ENTRYPOINT ["/app/entrypoint.sh"]
