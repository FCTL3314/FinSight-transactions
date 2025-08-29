from fastapi import FastAPI
from fastapi.responses import JSONResponse

from src.api.routers import base_router
from src.core import settings
from src.core.exceptions import DetailedHTTPException

app = FastAPI(
    title="FinSight Transactions",
    description="Backend server for managing financial transactions.",
    version="1.0.0",
)
app.include_router(base_router)


@app.exception_handler(DetailedHTTPException)
async def detailed_exception_handler(exc: DetailedHTTPException) -> JSONResponse:
    return JSONResponse(
        status_code=exc.status_code,
        content={"detail": exc.detail, "codename": exc.codename},
    )


if __name__ == "__main__":
    import uvicorn

    uvicorn.run(
        "main:app", host="0.0.0.0", port=settings.server.internal_port, reload=True
    )
