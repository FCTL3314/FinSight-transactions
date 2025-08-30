from typing import Callable

from pydantic import BaseModel, model_serializer


class IdFirstMixin(BaseModel):
    @model_serializer(mode="wrap")
    def _id_first(self, handler: Callable[[BaseModel], dict]) -> dict:
        data = handler(self)
        if "id" in data:
            return {"id": data["id"], **{k: v for k, v in data.items() if k != "id"}}
        return data

    class Config:
        arbitrary_types_allowed = True
