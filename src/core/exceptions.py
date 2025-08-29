from fastapi import HTTPException, status


class DetailedHTTPException(HTTPException):
    def __init__(self, status_code: int, detail: str, codename: str) -> None:
        super().__init__(status_code=status_code, detail=detail)
        self.codename = codename


ObjectNotFound = DetailedHTTPException(
    status_code=status.HTTP_404_NOT_FOUND,
    detail="Object not found",
    codename="not_found",
)

AccessDenied = DetailedHTTPException(
    status_code=status.HTTP_403_FORBIDDEN,
    detail="Access denied",
    codename="forbidden",
)
