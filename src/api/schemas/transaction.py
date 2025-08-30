from datetime import datetime, UTC

from fastapi.params import Query
from fastapi_pagination import LimitOffsetParams, LimitOffsetPage
from fastapi_pagination.customization import CustomizedPage, UseParams
from pydantic import BaseModel, Field

from src.api.schemas import IdFirstMixin
from src.core import settings


class TransactionBase(BaseModel):
    amount: float
    name: str
    note: str | None = None
    category_id: int | None = None
    made_at: datetime | None = Field(default_factory=lambda: datetime.now(UTC))


class TransactionCreate(TransactionBase):
    pass


class TransactionUpdate(BaseModel):
    amount: float | None = None
    name: str | None = None
    note: str | None = None
    category_id: int | None = None
    made_at: datetime | None = None


class TransactionResponse(IdFirstMixin, TransactionBase):
    id: int
    user_id: int
    created_at: datetime
    updated_at: datetime

    class Config:
        from_attributes = True


class TransactionPaginationParams(LimitOffsetParams):
    limit: int = Query(
        settings.pagination.transaction_limit,
        ge=1,
        le=settings.pagination.transaction_limit,
        description="Page size limit",
    )


TransactionPage = CustomizedPage[
    LimitOffsetPage[TransactionResponse],
    UseParams(TransactionPaginationParams),
]
