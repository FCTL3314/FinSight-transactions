from datetime import date
from typing import Sequence

from sqlalchemy import func, select
from sqlalchemy.ext.asyncio import AsyncSession

from src.db.models.detailing import FinanceDetailing
from src.db.models.transaction import Transaction


class DetailingRepository:
    def __init__(self, session: AsyncSession) -> None:
        self.session = session

    async def get_by_id(
        self, detailing_id: int, user_id: int
    ) -> FinanceDetailing | None:
        result = await self.session.execute(
            select(FinanceDetailing).filter(
                FinanceDetailing.id == detailing_id, FinanceDetailing.user_id == user_id
            )
        )
        return result.scalar_one_or_none()

    async def get_all(
        self, user_id: int, limit: int, offset: int
    ) -> tuple[Sequence[FinanceDetailing], int]:
        total_result = await self.session.execute(
            select(func.count()).where(FinanceDetailing.user_id == user_id)
        )
        total = total_result.scalar_one()

        result = await self.session.execute(
            select(FinanceDetailing)
            .where(FinanceDetailing.user_id == user_id)
            .order_by(FinanceDetailing.created_at.desc())
            .limit(limit)
            .offset(offset)
        )
        items = result.scalars().all()
        return items, total

    async def get_transaction_totals(
        self, user_id: int, date_from: date, date_to: date
    ) -> tuple[float, float]:
        income_query = await self.session.execute(
            select(func.coalesce(func.sum(Transaction.amount), 0.0)).filter(
                Transaction.user_id == user_id,
                Transaction.amount > 0,
                Transaction.made_at >= date_from,
                Transaction.made_at <= date_to,
            )
        )
        total_income = income_query.scalar()

        expense_query = await self.session.execute(
            select(func.coalesce(func.sum(Transaction.amount), 0.0)).filter(
                Transaction.user_id == user_id,
                Transaction.amount < 0,
                Transaction.made_at >= date_from,
                Transaction.made_at <= date_to,
            )
        )
        total_expense = expense_query.scalar()

        return total_income, abs(total_expense)

    async def create(self, detailing: FinanceDetailing) -> FinanceDetailing:
        self.session.add(detailing)
        await self.session.commit()
        await self.session.refresh(detailing)
        return detailing

    async def update(self, detailing: FinanceDetailing) -> FinanceDetailing:
        self.session.add(detailing)
        await self.session.commit()
        await self.session.refresh(detailing)
        return detailing

    async def delete(self, detailing: FinanceDetailing) -> None:
        await self.session.delete(detailing)
        await self.session.commit()
