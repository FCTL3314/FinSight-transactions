from fastapi import APIRouter, Depends, status
from fastapi_pagination import LimitOffsetPage

from src.api.dependencies import get_user_id, get_detailing_service
from src.api.schemas.detailing import (
    FinanceDetailingCreate,
    FinanceDetailingUpdate,
    FinanceDetailingResponse,
    FinanceDetailingPage,
    FinanceDetailingPaginationParams,
)
from src.services.detailing import DetailingService

router = APIRouter(prefix="/detailing", tags=["Detailing"])


@router.post(
    "/", response_model=FinanceDetailingResponse, status_code=status.HTTP_201_CREATED
)
async def create_detailing(
    detailing: FinanceDetailingCreate,
    user_id: int = Depends(get_user_id),
    service: DetailingService = Depends(get_detailing_service),
):
    return await service.create(detailing, user_id)


@router.get("/", response_model=FinanceDetailingPage)
async def get_all_detailing(
    user_id: int = Depends(get_user_id),
    params: FinanceDetailingPaginationParams = Depends(),
    service: DetailingService = Depends(get_detailing_service),
):
    items, total = await service.get_all(user_id, params.limit, params.offset)
    return LimitOffsetPage.create(items, params=params, total=total)


@router.get("/{detailing_id}", response_model=FinanceDetailingResponse)
async def get_detailing(
    detailing_id: int,
    user_id: int = Depends(get_user_id),
    service: DetailingService = Depends(get_detailing_service),
):
    return await service.get_by_id(detailing_id, user_id)


@router.patch("/{detailing_id}", response_model=FinanceDetailingResponse)
async def update_detailing(
    detailing_id: int,
    detailing_update: FinanceDetailingUpdate,
    user_id: int = Depends(get_user_id),
    service: DetailingService = Depends(get_detailing_service),
):
    return await service.update(detailing_id, detailing_update, user_id)


@router.delete("/{detailing_id}", status_code=status.HTTP_204_NO_CONTENT)
async def delete_detailing(
    detailing_id: int,
    user_id: int = Depends(get_user_id),
    service: DetailingService = Depends(get_detailing_service),
):
    await service.delete(detailing_id, user_id)
