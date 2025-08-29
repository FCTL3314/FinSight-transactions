from datetime import date

from sqlalchemy import func
from sqlalchemy.orm import Session

from src.db.models.detailing import FinanceDetailing
from src.db.models.transaction import Transaction


class DetailingRepository:
    def __init__(self, db: Session):
        self.db = db

    def get_by_id(self, detailing_id: int, user_id: int) -> FinanceDetailing | None:
        return (
            self.db.query(FinanceDetailing)
            .filter(
                FinanceDetailing.id == detailing_id, FinanceDetailing.user_id == user_id
            )
            .first()
        )

    def get_all(
        self, user_id: int, skip: int = 0, limit: int = 32
    ) -> list[type[FinanceDetailing]]:
        return (
            self.db.query(FinanceDetailing)
            .filter(FinanceDetailing.user_id == user_id)
            .order_by(FinanceDetailing.created_at.desc())
            .offset(skip)
            .limit(limit)
            .all()
        )

    def get_transaction_totals(
        self, user_id: int, date_from: date, date_to: date
    ) -> tuple[float, float]:
        income_query = (
            self.db.query(func.coalesce(func.sum(Transaction.amount), 0.0))
            .filter(
                Transaction.user_id == user_id,
                Transaction.amount > 0,
                Transaction.made_at >= date_from,
                Transaction.made_at <= date_to,
            )
            .scalar()
        )

        expense_query = (
            self.db.query(func.coalesce(func.sum(Transaction.amount), 0.0))
            .filter(
                Transaction.user_id == user_id,
                Transaction.amount < 0,
                Transaction.made_at >= date_from,
                Transaction.made_at <= date_to,
            )
            .scalar()
        )

        return income_query, abs(expense_query)

    def create(self, detailing: FinanceDetailing) -> FinanceDetailing:
        self.db.add(detailing)
        self.db.commit()
        self.db.refresh(detailing)
        return detailing

    def update(self, detailing: FinanceDetailing) -> FinanceDetailing:
        self.db.add(detailing)
        self.db.commit()
        self.db.refresh(detailing)
        return detailing

    def delete(self, detailing: FinanceDetailing) -> None:
        self.db.delete(detailing)
        self.db.commit()
