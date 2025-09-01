import logging.config

from src.core import settings

LOGGING_CONFIG_DICT = {
    "version": 1,
    "disable_existing_loggers": False,
    "formatters": {
        "default": {
            "format": "%(asctime)s | %(levelname)-8s | %(name)s:%(funcName)s:%(lineno)d | %(message)s",
        },
    },
    "handlers": {
        "console": {
            "level": "DEBUG",
            "class": "logging.StreamHandler",
            "formatter": "default",
            "stream": "ext://sys.stdout",
        },
        "app_file": {
            "level": "INFO",
            "class": "logging.handlers.RotatingFileHandler",
            "formatter": "default",
            "filename": settings.logs_dir / "app.log",
            "maxBytes": 10485760,  # 10MB
            "backupCount": 5,
            "encoding": "utf8",
        },
        "services_file": {
            "level": "DEBUG",
            "class": "logging.handlers.RotatingFileHandler",
            "formatter": "default",
            "filename": settings.logs_dir / "services.log",
            "maxBytes": 10485760,
            "backupCount": 5,
            "encoding": "utf8",
        },
        "repositories_file": {
            "level": "DEBUG",
            "class": "logging.handlers.RotatingFileHandler",
            "formatter": "default",
            "filename": settings.logs_dir / "repositories.log",
            "maxBytes": 10485760,
            "backupCount": 5,
            "encoding": "utf8",
        },
        "routers_file": {
            "level": "DEBUG",
            "class": "logging.handlers.RotatingFileHandler",
            "formatter": "default",
            "filename": settings.logs_dir / "routers.log",
            "maxBytes": 10485760,
            "backupCount": 5,
            "encoding": "utf8",
        },
    },
    "loggers": {
        "src.services": {
            "level": "DEBUG",
            "handlers": ["services_file", "console"],
            "propagate": False,
        },
        "src.repositories": {
            "level": "DEBUG",
            "handlers": ["repositories_file", "console"],
            "propagate": False,
        },
        "src.api.routers": {
            "level": "DEBUG",
            "handlers": ["routers_file", "console"],
            "propagate": False,
        },
    },
    "root": {
        "level": "INFO",
        "handlers": ["app_file", "console"],
    },
}


def setup_logging() -> None:
    settings.logs_dir.mkdir(exist_ok=True)
    logging.config.dictConfig(LOGGING_CONFIG_DICT)
