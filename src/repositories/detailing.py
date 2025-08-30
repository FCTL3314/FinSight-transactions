from datetime import date
from sqlalchemy import func, select
from sqlalchemy.ext.asyncio import AsyncSession
from src.db.models.detailing import FinanceDetailing
from src.db.models.transaction import Transaction


class DetailingRepository:
    def __init__(self, db: AsyncSession) -> None:
        self.db = db

    async def get_by_id(
        self, detailing_id: int, user_id: int
    ) -> FinanceDetailing | None:
        result = await self.db.execute(
            select(FinanceDetailing).filter(
                FinanceDetailing.id == detailing_id, FinanceDetailing.user_id == user_id
            )
        )
        return result.scalar_one_or_none()

    async def get_all(
        self, user_id: int, skip: int = 0, limit: int = 32
    ) -> list[FinanceDetailing]:
        result = await self.db.execute(
            select(FinanceDetailing)
            .filter(FinanceDetailing.user_id == user_id)
            .order_by(FinanceDetailing.created_at.desc())
            .offset(skip)
            .limit(limit)
        )
        return result.scalars().all()

    async def get_transaction_totals(
        self, user_id: int, date_from: date, date_to: date
    ) -> tuple[float, float]:
        income_query = await self.db.execute(
            select(func.coalesce(func.sum(Transaction.amount), 0.0)).filter(
                Transaction.user_id == user_id,
                Transaction.amount > 0,
                Transaction.made_at >= date_from,
                Transaction.made_at <= date_to,
            )
        )
        total_income = income_query.scalar()

        expense_query = await self.db.execute(
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
        self.db.add(detailing)
        await self.db.commit()
        await self.db.refresh(detailing)
        return detailing

    async def update(self, detailing: FinanceDetailing) -> FinanceDetailing:
        self.db.add(detailing)
        await self.db.commit()
        await self.db.refresh(detailing)
        return detailing

    async def delete(self, detailing: FinanceDetailing) -> None:
        await self.db.delete(detailing)
        await self.db.commit()
