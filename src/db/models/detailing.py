from sqlalchemy import Column, Integer, Float, DateTime, BigInteger, Date
from sqlalchemy.sql import func

from src.db.models import Base


class FinanceDetailing(Base):
    __tablename__ = "finance_detailing"

    id = Column(Integer, primary_key=True, index=True)
    user_id = Column(BigInteger, nullable=False)
    date_from = Column(Date, nullable=False)
    date_to = Column(Date, nullable=False)
    initial_amount = Column(Float, nullable=False)
    current_amount = Column(Float, nullable=False)
    total_income = Column(Float, default=0.0)
    total_expense = Column(Float, default=0.0)
    profit_estimated = Column(Float, default=0.0)
    profit_real = Column(Float, default=0.0)
    after_amount_net = Column(Float, default=0.0)
    after_amount_gross = Column(Float, default=0.0)
    created_at = Column(DateTime(timezone=True), server_default=func.now())
    updated_at = Column(
        DateTime(timezone=True), default=func.now(), onupdate=func.now()
    )
