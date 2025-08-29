FROM python:3.11-slim

RUN pip install "uvicorn[standard]" fastapi pydantic pydantic-settings sqlalchemy psycopg2-binary alembic python-dotenv pyyaml

RUN pip install watchgod

WORKDIR /app

COPY ./requirements.txt /app/requirements.txt
RUN pip install --no-cache-dir -r /app/requirements.txt

COPY . /app

EXPOSE ${INTERNAL_PORT}

ENTRYPOINT ["/app/docker/local/scripts/entrypoint.sh"]
