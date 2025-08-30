from typing import Sequence
from sqlalchemy.ext.asyncio import AsyncSession

from src.api.schemas.detailing import FinanceDetailingCreate, FinanceDetailingUpdate
from src.core.exceptions import ObjectNotFound
from src.db.models.detailing import FinanceDetailing
from src.repositories.detailing import DetailingRepository


class DetailingService:
    def __init__(self, db: AsyncSession) -> None:
        self.db = db
        self.repo = DetailingRepository(self.db)

    async def _calculate_and_set_fields(self, detailing: FinanceDetailing) -> None:
        total_income, total_expense = await self.repo.get_transaction_totals(
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

    async def get_by_id(self, detailing_id: int, user_id: int) -> FinanceDetailing:
        detailing = await self.repo.get_by_id(detailing_id, user_id)
        if not detailing:
            raise ObjectNotFound
        return detailing

    async def get_all(
        self, user_id: int, limit: int, offset: int
    ) -> tuple[Sequence[FinanceDetailing], int]:
        return await self.repo.get_all(user_id, limit, offset)

    async def create(
        self, detailing_data: FinanceDetailingCreate, user_id: int
    ) -> FinanceDetailing:
        db_detailing = FinanceDetailing(**detailing_data.model_dump(), user_id=user_id)
        await self._calculate_and_set_fields(db_detailing)
        return await self.repo.create(db_detailing)

    async def update(
        self, detailing_id: int, detailing_update: FinanceDetailingUpdate, user_id: int
    ) -> FinanceDetailing:
        db_detailing = await self.get_by_id(detailing_id, user_id)
        update_data = detailing_update.model_dump(exclude_unset=True)
        for key, value in update_data.items():
            setattr(db_detailing, key, value)
        await self._calculate_and_set_fields(db_detailing)
        return await self.repo.update(db_detailing)

    async def delete(self, detailing_id: int, user_id: int) -> None:
        detailing = await self.get_by_id(detailing_id, user_id)
        await self.repo.delete(detailing)
