from typing import Sequence

from sqlalchemy import select, delete, update
from sqlalchemy.ext.asyncio import AsyncSession

from src.api.schemas.transaction import TransactionCreate, TransactionUpdate
from src.db.models.transaction import Transaction


class TransactionRepository:
    def __init__(self, db: AsyncSession) -> None:
        self.db = db

    async def get_by_id(self, transaction_id: int, user_id: int) -> Transaction | None:
        result = await self.db.execute(
            select(Transaction).filter(
                Transaction.id == transaction_id, Transaction.user_id == user_id
            )
        )
        return result.scalar_one_or_none()

    async def get_all(
        self, user_id: int, skip: int = 0, limit: int = 100
    ) -> Sequence[Transaction]:
        result = await self.db.execute(
            select(Transaction)
            .filter(Transaction.user_id == user_id)
            .offset(skip)
            .limit(limit)
        )
        return result.scalars().all()

    async def create(self, transaction: TransactionCreate, user_id: int) -> Transaction:
        db_transaction = Transaction(**transaction.model_dump(), user_id=user_id)
        self.db.add(db_transaction)
        await self.db.commit()
        await self.db.refresh(db_transaction)
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
        result = await self.db.execute(stmt)
        updated_transaction = result.scalar_one_or_none()
        if updated_transaction:
            await self.db.commit()
            await self.db.refresh(updated_transaction)
        return updated_transaction

    async def delete(self, transaction_id: int, user_id: int) -> None:
        stmt = delete(Transaction).where(
            Transaction.id == transaction_id, Transaction.user_id == user_id
        )
        await self.db.execute(stmt)
        await self.db.commit()
