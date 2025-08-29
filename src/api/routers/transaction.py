from fastapi import APIRouter, Depends, status

from src.api.dependencies import get_user_id, get_transactions_service
from src.api.schemas.transaction import (
    TransactionCreate,
    TransactionUpdate,
    TransactionResponse,
)
from src.services.transaction import TransactionService

router = APIRouter(prefix="/transactions", tags=["Transactions"])


@router.post(
    "/", response_model=TransactionResponse, status_code=status.HTTP_201_CREATED
)
def create_transaction(
    transaction: TransactionCreate,
    user_id: int = Depends(get_user_id),
    service: TransactionService = Depends(get_transactions_service),
):
    return service.create(transaction, user_id)


@router.get("/{transaction_id}", response_model=TransactionResponse)
def get_transaction(
    transaction_id: int,
    user_id: int = Depends(get_user_id),
    service: TransactionService = Depends(get_transactions_service),
):
    return service.get_by_id(transaction_id, user_id)


@router.get("/", response_model=list[TransactionResponse])
def get_all_transactions(
    user_id: int = Depends(get_user_id),
    skip: int = 0,
    limit: int = 100,
    service: TransactionService = Depends(get_transactions_service),
):
    return service.get_all(user_id, skip, limit)


@router.patch("/{transaction_id}", response_model=TransactionResponse)
def update_transaction(
    transaction_id: int,
    transaction: TransactionUpdate,
    user_id: int = Depends(get_user_id),
    service: TransactionService = Depends(get_transactions_service),
):
    return service.update(transaction_id, transaction, user_id)


@router.delete("/{transaction_id}", status_code=status.HTTP_204_NO_CONTENT)
def delete_transaction(
    transaction_id: int,
    user_id: int = Depends(get_user_id),
    service: TransactionService = Depends(get_transactions_service),
):
    service.delete(transaction_id, user_id)
