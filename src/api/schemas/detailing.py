from datetime import date, datetime

from pydantic import BaseModel, model_serializer

from src.api.schemas import IdFirstMixin


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
