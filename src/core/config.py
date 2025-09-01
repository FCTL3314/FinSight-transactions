from pathlib import Path

import yaml
from pydantic_settings import BaseSettings, SettingsConfigDict

from src.core.types import LoggingLevels


class PaginationSettings(BaseSettings):
    transaction_limit: int
    finance_detailing_limit: int


class ServerSettings(BaseSettings):
    debug: bool
    internal_port: int
    external_port: int

    model_config = SettingsConfigDict(extra="ignore")


class DatabaseSettings(BaseSettings):
    name: str
    user: str
    password: str
    host: str
    port: int

    model_config = SettingsConfigDict(env_prefix="DB_", extra="ignore")

    @property
    def url(self) -> str:
        return (
            f"postgresql+asyncpg://{self.user}:{self.password}@"
            f"{self.host}:{self.port}/{self.name}"
        )


class Settings(BaseSettings):
    base_dir: Path = Path(__file__).resolve().parent.parent.parent
    logs_dir: Path
    settings_dir: Path = base_dir / "settings"

    log_level: LoggingLevels

    server: ServerSettings = ServerSettings()  # noqa
    db: DatabaseSettings = DatabaseSettings()  # noqa

    config_file_path: Path = settings_dir / "config.yml"
    _config_file_dict = yaml.safe_load(config_file_path.read_text())

    pagination: PaginationSettings = PaginationSettings(
        **_config_file_dict["pagination"]
    )
