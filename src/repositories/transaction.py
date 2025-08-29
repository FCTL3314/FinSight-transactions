from sqlalchemy.orm import Session

from src.api.schemas.transaction import TransactionCreate, TransactionUpdate
from src.db.models.transaction import Transaction


class TransactionRepository:
    def __init__(self, db: Session):
        self.db = db

    def get_by_id(self, transaction_id: int, user_id: int) -> Transaction | None:
        return (
            self.db.query(Transaction)
            .filter(Transaction.id == transaction_id, Transaction.user_id == user_id)
            .first()
        )

    def get_all(
        self, user_id: int, skip: int = 0, limit: int = 100
    ) -> list[type[Transaction]]:
        return (
            self.db.query(Transaction)
            .filter(Transaction.user_id == user_id)
            .offset(skip)
            .limit(limit)
            .all()
        )

    def create(self, transaction: TransactionCreate, user_id: int) -> Transaction:
        db_transaction = Transaction(**transaction.model_dump(), user_id=user_id)
        self.db.add(db_transaction)
        self.db.commit()
        self.db.refresh(db_transaction)
        return db_transaction

    def update(
        self, transaction_id: int, transaction: TransactionUpdate, user_id: int
    ) -> Transaction | None:
        db_transaction = self.get_by_id(transaction_id, user_id)
        if db_transaction:
            update_data = transaction.model_dump(exclude_unset=True)
            for key, value in update_data.items():
                setattr(db_transaction, key, value)
            self.db.commit()
            self.db.refresh(db_transaction)
        return db_transaction

    def delete(self, transaction_id: int, user_id: int) -> None:
        db_transaction = self.get_by_id(transaction_id, user_id)
        if db_transaction:
            self.db.delete(db_transaction)
            self.db.commit()
