from fastapi import APIRouter
from src.api.routers.system import router as system_router
from src.api.routers.transaction import router as transaction_router
from src.api.routers.detailing import router as detailing_router

base_router = APIRouter()
base_router.include_router(system_router)

v1_router = APIRouter(prefix="/api/v1")
v1_router.include_router(transaction_router)
v1_router.include_router(detailing_router)

base_router.include_router(v1_router)
