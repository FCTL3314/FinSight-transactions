import logging.config
from src.core import settings


def setup_logging() -> None:
    settings.logs_dir.mkdir(exist_ok=True)
    logging.config.dictConfig(settings.logging_config_dict)
