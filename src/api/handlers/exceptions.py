import logging
from http import HTTPStatus

from fastapi.responses import JSONResponse
from starlette.requests import Request

from src.core.exceptions import DetailedHTTPException

logger = logging.getLogger(__name__)


async def detailed_http_exception_handler(_: Request, exc: Exception) -> JSONResponse:
    if not isinstance(exc, DetailedHTTPException):
        logger.error(f"An unexpected exception occurred: {exc}", exc_info=True)
        return JSONResponse(
            status_code=HTTPStatus.INTERNAL_SERVER_ERROR,
            content={"detail": "Internal Server Error", "codename": "internal_error"},
        )

    logger.error(f"A DetailedHTTPException occurred: {exc.detail}", exc_info=True)
    return JSONResponse(
        status_code=exc.status_code,
        content={"detail": exc.detail, "codename": exc.codename},
    )
