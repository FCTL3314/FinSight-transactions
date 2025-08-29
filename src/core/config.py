from pathlib import Path

import yaml
from pydantic_settings import BaseSettings, SettingsConfigDict

_BASE_DIR: Path = Path(__file__).resolve().parent.parent.parent
_SETTINGS_DIR: Path = _BASE_DIR / "settings"
_LOCAL_ENV_PATH = _SETTINGS_DIR / ".env.local"
_CONFIG_FILE_PATH = _SETTINGS_DIR / "config.yml"


class PaginationSettings(BaseSettings):
    transaction_limit: int
    finance_detailing_limit: int


class ServerSettings(BaseSettings):
    debug: bool
    internal_port: int
    external_port: int

    model_config = SettingsConfigDict(env_file=_LOCAL_ENV_PATH, extra="ignore")


class DatabaseSettings(BaseSettings):
    name: str
    user: str
    password: str
    host: str
    port: int

    model_config = SettingsConfigDict(
        env_file=_LOCAL_ENV_PATH, env_prefix="DB_", extra="ignore"
    )

    @property
    def url(self) -> str:
        return (
            f"postgresql+psycopg2://{self.user}:{self.password}@"
            f"{self.host}:{self.port}/{self.name}"
        )


class Settings(BaseSettings):
    base_dir: Path = _BASE_DIR
    settings_dir: Path = _SETTINGS_DIR

    server: ServerSettings = ServerSettings()  # noqa
    db: DatabaseSettings = DatabaseSettings()  # noqa

    _config_file_dict = yaml.safe_load(_CONFIG_FILE_PATH.read_text())

    pagination: PaginationSettings = PaginationSettings(
        **_config_file_dict["pagination"]
    )
