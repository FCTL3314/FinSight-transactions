from fastapi import FastAPI

from src.api.handlers.lifespan import lifespan
from src.core import settings

app = FastAPI(
    title="FinSight Transactions",
    description="Backend server for managing financial transactions.",
    version="1.0.0",
    lifespan=lifespan,
    debug=settings.server.debug,
)


if __name__ == "__main__":
    import uvicorn

    uvicorn.run(
        "main:app",
        host="0.0.0.0",
        port=settings.server.internal_port,
        reload=True,
    )
