import logging
import sys
from logging.handlers import RotatingFileHandler

from src.core import settings


def setup_logging() -> None:
    settings.logs_dir.mkdir(exist_ok=True)
    log_file = settings.logs_dir / "app.log"

    root_logger = logging.getLogger()
    root_logger.setLevel(settings.log_level)

    formatter = logging.Formatter(
        "%(asctime)s | %(levelname)-8s | %(name)s:%(funcName)s:%(lineno)d | %(message)s"
    )

    console_handler = logging.StreamHandler(sys.stdout)
    console_handler.setFormatter(formatter)
    root_logger.addHandler(console_handler)

    file_handler = RotatingFileHandler(
        log_file, maxBytes=1024 * 1024 * 5, backupCount=5
    )
    file_handler.setFormatter(formatter)
    root_logger.addHandler(file_handler)
