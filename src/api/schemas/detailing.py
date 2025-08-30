from datetime import date, datetime

from fastapi.params import Query
from fastapi_pagination import LimitOffsetParams, LimitOffsetPage
from fastapi_pagination.customization import CustomizedPage, UseParams
from pydantic import BaseModel

from src.api.schemas import IdFirstMixin
from src.core import settings


class FinanceDetailingBase(BaseModel):
    date_from: date
    date_to: date
    initial_amount: float
    current_amount: float


class FinanceDetailingCreate(FinanceDetailingBase):
    pass


class FinanceDetailingUpdate(BaseModel):
    date_from: date | None = None
    date_to: date | None = None
    initial_amount: float | None = None
    current_amount: float | None = None


class FinanceDetailingResponse(IdFirstMixin, FinanceDetailingBase):
    id: int
    user_id: int
    total_income: float
    total_expense: float
    profit_estimated: float
    profit_real: float
    after_amount_net: float
    after_amount_gross: float
    created_at: datetime
    updated_at: datetime

    class Config:
        from_attributes = True


class FinanceDetailingPaginationParams(LimitOffsetParams):
    limit: int = Query(
        settings.pagination.finance_detailing_limit,
        ge=1,
        le=settings.pagination.finance_detailing_limit,
        description="Page size limit",
    )


FinanceDetailingPage = CustomizedPage[
    LimitOffsetPage[FinanceDetailingResponse],
    UseParams(FinanceDetailingPaginationParams),
]
