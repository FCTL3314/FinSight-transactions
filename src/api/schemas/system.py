from typing import Literal

from pydantic import BaseModel


class HealthCheckResponse(BaseModel):
    status: Literal["ok", "error"]
    db_connection: bool
