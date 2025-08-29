FROM python:3.13-slim as builder

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

RUN pip install uv

WORKDIR /app

COPY pyproject.toml uv.lock ./

RUN uv pip sync --no-dev --system

COPY . .


FROM python:3.13-slim

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1
ENV APP_USER=appuser

RUN groupadd -r ${APP_USER} && useradd --no-log-init -r -g ${APP_USER} ${APP_USER}

WORKDIR /app

COPY --from=builder /usr/local/lib/python3.13/site-packages /usr/local/lib/python3.13/site-packages
COPY --from=builder /app .

COPY ./docker/prod/scripts/entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh

RUN chown -R ${APP_USER}:${APP_USER} /app

USER ${APP_USER}

EXPOSE ${INTERNAL_PORT}

ENTRYPOINT ["/app/entrypoint.sh"]
