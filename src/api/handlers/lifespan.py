import logging
from contextlib import asynccontextmanager
from typing import AsyncGenerator

from fastapi import FastAPI

from src.api.handlers.exceptions import detailed_http_exception_handler
from src.api.routers import base_router
from src.core.exceptions import DetailedHTTPException
from src.core.logging import setup_logging

logger = logging.getLogger(__name__)


@asynccontextmanager
async def lifespan(app: FastAPI) -> AsyncGenerator[None, None]:
    setup_logging()
    app.include_router(base_router)
    app.add_exception_handler(DetailedHTTPException, detailed_http_exception_handler)
    logger.info("Application startup completed.")
    yield
