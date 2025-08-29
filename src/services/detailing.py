from sqlalchemy.orm import Session
from src.repositories.detailing import DetailingRepository
from src.api.schemas.detailing import FinanceDetailingCreate, FinanceDetailingUpdate
from src.core.exceptions import ObjectNotFound
from src.db.models.detailing import FinanceDetailing


class DetailingService:
    def __init__(self, db: Session):
        self.db = db
        self.repo = DetailingRepository(self.db)

    def _calculate_and_set_fields(self, detailing: FinanceDetailing) -> None:
        total_income, total_expense = self.repo.get_transaction_totals(
            user_id=detailing.user_id,
            date_from=detailing.date_from,
            date_to=detailing.date_to,
        )

        detailing.total_income = total_income
        detailing.total_expense = total_expense

        detailing.profit_estimated = detailing.total_income - detailing.total_expense
        detailing.profit_real = detailing.current_amount - detailing.initial_amount
        detailing.after_amount_gross = detailing.initial_amount + detailing.total_income
        detailing.after_amount_net = (
            detailing.after_amount_gross - detailing.total_expense
        )

    def get_by_id(self, detailing_id: int, user_id: int) -> FinanceDetailing:
        detailing = self.repo.get_by_id(detailing_id, user_id)
        if not detailing:
            raise ObjectNotFound
        return detailing

    def get_all(self, user_id: int, skip: int, limit: int) -> list[type[FinanceDetailing]]:
        return self.repo.get_all(user_id, skip, limit)

    def create(self, detailing_data: FinanceDetailingCreate, user_id: int) -> FinanceDetailing:
        db_detailing = FinanceDetailing(**detailing_data.model_dump(), user_id=user_id)

        self._calculate_and_set_fields(db_detailing)

        return self.repo.create(db_detailing)

    def update(
        self, detailing_id: int, detailing_update: FinanceDetailingUpdate, user_id: int
    ) -> FinanceDetailing:
        db_detailing = self.get_by_id(detailing_id, user_id)

        update_data = detailing_update.model_dump(exclude_unset=True)
        for key, value in update_data.items():
            setattr(db_detailing, key, value)

        self._calculate_and_set_fields(db_detailing)

        return self.repo.update(db_detailing)

    def delete(self, detailing_id: int, user_id: int) -> None:
        detailing = self.get_by_id(detailing_id, user_id)
        self.repo.delete(detailing)
