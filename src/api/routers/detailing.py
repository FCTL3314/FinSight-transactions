from fastapi import APIRouter, Depends, status

from src.api.dependencies import get_user_id, get_detailing_service
from src.api.schemas.detailing import (
    FinanceDetailingCreate,
    FinanceDetailingUpdate,
    FinanceDetailingResponse,
)
from src.core import settings
from src.services.detailing import DetailingService

router = APIRouter(prefix="/detailing", tags=["Detailing"])


@router.post(
    "/", response_model=FinanceDetailingResponse, status_code=status.HTTP_201_CREATED
)
def create_detailing(
    detailing: FinanceDetailingCreate,
    user_id: int = Depends(get_user_id),
    service: DetailingService = Depends(get_detailing_service),
):
    return service.create(detailing, user_id)


@router.get("/", response_model=list[FinanceDetailingResponse])
def get_all_detailing(
    user_id: int = Depends(get_user_id),
    skip: int = 0,
    limit: int = settings.pagination.finance_detailing_limit,
    service: DetailingService = Depends(get_detailing_service),
):
    return service.get_all(user_id, skip, limit)


@router.get("/{detailing_id}", response_model=FinanceDetailingResponse)
def get_detailing(
    detailing_id: int,
    user_id: int = Depends(get_user_id),
    service: DetailingService = Depends(get_detailing_service),
):
    return service.get_by_id(detailing_id, user_id)


@router.patch("/{detailing_id}", response_model=FinanceDetailingResponse)
def update_detailing(
    detailing_id: int,
    detailing_update: FinanceDetailingUpdate,
    user_id: int = Depends(get_user_id),
    service: DetailingService = Depends(get_detailing_service),
):
    return service.update(detailing_id, detailing_update, user_id)


@router.delete("/{detailing_id}", status_code=status.HTTP_204_NO_CONTENT)
def delete_detailing(
    detailing_id: int,
    user_id: int = Depends(get_user_id),
    service: DetailingService = Depends(get_detailing_service),
):
    service.delete(detailing_id, user_id)
