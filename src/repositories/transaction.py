from typing import Sequence

from sqlalchemy import select, delete, update, func
from sqlalchemy.ext.asyncio import AsyncSession

from src.api.schemas.transaction import TransactionCreate, TransactionUpdate
from src.db.models.transaction import Transaction


class TransactionRepository:
    def __init__(self, session: AsyncSession) -> None:
        self.session = session

    async def get_by_id(self, transaction_id: int, user_id: int) -> Transaction | None:
        result = await self.session.execute(
            select(Transaction).filter(
                Transaction.id == transaction_id, Transaction.user_id == user_id
            )
        )
        return result.scalar_one_or_none()

    async def get_all(
        self, user_id: int, limit: int, offset: int
    ) -> tuple[Sequence[Transaction], int]:
        total_result = await self.session.execute(
            select(func.count()).where(Transaction.user_id == user_id)
        )
        total = total_result.scalar_one()

        result = await self.session.execute(
            select(Transaction)
            .where(Transaction.user_id == user_id)
            .order_by(Transaction.created_at.desc())
            .limit(limit)
            .offset(offset)
        )
        items = result.scalars().all()
        return items, total

    async def create(self, transaction: TransactionCreate, user_id: int) -> Transaction:
        db_transaction = Transaction(**transaction.model_dump(), user_id=user_id)
        self.session.add(db_transaction)
        await self.session.commit()
        await self.session.refresh(db_transaction)
        return db_transaction

    async def update(
        self, transaction_id: int, transaction: TransactionUpdate, user_id: int
    ) -> Transaction | None:
        stmt = (
            update(Transaction)
            .where(Transaction.id == transaction_id, Transaction.user_id == user_id)
            .values(**transaction.model_dump(exclude_unset=True))
            .returning(Transaction)
        )
        result = await self.session.execute(stmt)
        updated_transaction = result.scalar_one_or_none()
        if updated_transaction:
            await self.session.commit()
            await self.session.refresh(updated_transaction)
        return updated_transaction

    async def delete(self, transaction_id: int, user_id: int) -> None:
        stmt = delete(Transaction).where(
            Transaction.id == transaction_id, Transaction.user_id == user_id
        )
        await self.session.execute(stmt)
        await self.session.commit()
