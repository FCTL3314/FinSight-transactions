from sqlalchemy.orm import Session

from src.api.schemas.transaction import TransactionCreate, TransactionUpdate
from src.core.exceptions import ObjectNotFound
from src.db.models.transaction import Transaction
from src.repositories.transaction import TransactionRepository


class TransactionService:
    def __init__(self, db: Session):
        self.repo = TransactionRepository(db)

    def get_by_id(self, transaction_id: int, user_id: int) -> Transaction:
        transaction = self.repo.get_by_id(transaction_id, user_id)
        if not transaction:
            raise ObjectNotFound
        return transaction

    def get_all(self, user_id: int, skip: int, limit: int) -> list[type[Transaction]]:
        return self.repo.get_all(user_id, skip, limit)

    def create(self, transaction: TransactionCreate, user_id: int) -> Transaction:
        return self.repo.create(transaction, user_id)

    def update(self, transaction_id: int, transaction: TransactionUpdate, user_id: int) -> Transaction:
        updated_transaction = self.repo.update(transaction_id, transaction, user_id)
        if not updated_transaction:
            raise ObjectNotFound
        return updated_transaction

    def delete(self, transaction_id: int, user_id: int) -> None:
        self.repo.delete(transaction_id, user_id)
