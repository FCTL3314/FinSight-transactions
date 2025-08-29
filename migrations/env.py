from logging.config import fileConfig

from alembic import context
from src.core import settings
from src.db.database import engine
from src.db.models import Base
from src.db.models.detailing import FinanceDetailing  # noqa
from src.db.models.transaction import Transaction  # noqa

config = context.config
section = config.config_ini_section

config.set_section_option(section, "DB_NAME", settings.db.name)
config.set_section_option(section, "DB_USER", settings.db.user)
config.set_section_option(section, "DB_PASSWORD", settings.db.password)
config.set_section_option(section, "DB_HOST", settings.db.host)
config.set_section_option(section, "DB_PORT", str(settings.db.port))

if config.config_file_name is not None:
    fileConfig(config.config_file_name)

target_metadata = Base.metadata


def run_migrations_offline() -> None:
    url = settings.db.url
    context.configure(
        url=url,
        target_metadata=target_metadata,
        literal_binds=True,
        dialect_opts={"paramstyle": "named"},
    )

    with context.begin_transaction():
        context.run_migrations()


def run_migrations_online() -> None:
    with engine.connect() as connection:
        context.configure(connection=connection, target_metadata=target_metadata)

        with context.begin_transaction():
            context.run_migrations()


if context.is_offline_mode():
    run_migrations_offline()
else:
    run_migrations_online()
