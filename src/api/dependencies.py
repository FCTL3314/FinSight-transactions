from http import HTTPStatus
from typing import Annotated

from fastapi import Header, HTTPException, Depends
from sqlalchemy.ext.asyncio import AsyncSession

from src.db.database import get_session
from src.services.detailing import DetailingService
from src.services.transaction import TransactionService


async def get_user_id(x_user_id: Annotated[int | None, Header()] = None) -> int:
    if x_user_id is None:
        raise HTTPException(
            status_code=HTTPStatus.UNAUTHORIZED, detail="X-User-ID header is missing"
        )
    return x_user_id


async def get_transactions_service(
    session: AsyncSession = Depends(get_session),
) -> TransactionService:
    return TransactionService(session)


async def get_detailing_service(
    session: AsyncSession = Depends(get_session),
) -> DetailingService:
    return DetailingService(session)
