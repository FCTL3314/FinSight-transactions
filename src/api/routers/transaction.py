from fastapi import APIRouter, Depends, status
from fastapi_pagination import LimitOffsetPage

from src.api.dependencies import get_user_id, get_transactions_service
from src.api.schemas.transaction import (
    TransactionCreate,
    TransactionUpdate,
    TransactionResponse,
    TransactionPage,
    TransactionPaginationParams,
)
from src.services.transaction import TransactionService

router = APIRouter(prefix="/transactions", tags=["Transactions"])


@router.post(
    "/", response_model=TransactionResponse, status_code=status.HTTP_201_CREATED
)
async def create_transaction(
    transaction: TransactionCreate,
    user_id: int = Depends(get_user_id),
    service: TransactionService = Depends(get_transactions_service),
):
    return await service.create(transaction, user_id)


@router.get("/{transaction_id}", response_model=TransactionResponse)
async def get_transaction(
    transaction_id: int,
    user_id: int = Depends(get_user_id),
    service: TransactionService = Depends(get_transactions_service),
):
    return await service.get_by_id(transaction_id, user_id)


@router.get("/", response_model=TransactionPage)
async def get_all_transactions(
    user_id: int = Depends(get_user_id),
    params: TransactionPaginationParams = Depends(),
    service: TransactionService = Depends(get_transactions_service),
):
    items, total = await service.get_all(user_id, params.limit, params.offset)
    return LimitOffsetPage.create(items, params=params, total=total)


@router.patch("/{transaction_id}", response_model=TransactionResponse)
async def update_transaction(
    transaction_id: int,
    transaction: TransactionUpdate,
    user_id: int = Depends(get_user_id),
    service: TransactionService = Depends(get_transactions_service),
):
    return await service.update(transaction_id, transaction, user_id)


@router.delete("/{transaction_id}", status_code=status.HTTP_204_NO_CONTENT)
async def delete_transaction(
    transaction_id: int,
    user_id: int = Depends(get_user_id),
    service: TransactionService = Depends(get_transactions_service),
):
    await service.delete(transaction_id, user_id)
