import logging.config
from src.core import settings


def setup_logging() -> None:
    logging.config.dictConfig(settings.logging_config_dict)
