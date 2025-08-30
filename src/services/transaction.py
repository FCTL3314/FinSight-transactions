from typing import Sequence

from sqlalchemy.ext.asyncio import AsyncSession

from src.api.schemas.transaction import TransactionCreate, TransactionUpdate
from src.core.exceptions import ObjectNotFound
from src.db.models.transaction import Transaction
from src.repositories.transaction import TransactionRepository


class TransactionService:
    def __init__(self, db: AsyncSession) -> None:
        self.repo = TransactionRepository(db)

    async def get_by_id(self, transaction_id: int, user_id: int) -> Transaction:
        transaction = await self.repo.get_by_id(transaction_id, user_id)
        if not transaction:
            raise ObjectNotFound
        return transaction

    async def get_all(
        self, user_id: int, limit: int, offset: int
    ) -> tuple[Sequence[Transaction], int]:
        return await self.repo.get_all(user_id, limit, offset)

    async def create(self, transaction: TransactionCreate, user_id: int) -> Transaction:
        return await self.repo.create(transaction, user_id)

    async def update(
        self, transaction_id: int, transaction: TransactionUpdate, user_id: int
    ) -> Transaction:
        updated_transaction = await self.repo.update(
            transaction_id, transaction, user_id
        )
        if not updated_transaction:
            raise ObjectNotFound
        return updated_transaction

    async def delete(self, transaction_id: int, user_id: int) -> None:
        await self.repo.delete(transaction_id, user_id)
