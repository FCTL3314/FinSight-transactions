from logging.config import fileConfig

from alembic import context
from src.core import settings
from src.db.database import engine
from src.db.models import Base
from src.db.models.detailing import FinanceDetailing
from src.db.models.transaction import Transaction

# this is the Alembic Config object, which provides
# access to the values within the .ini file in use.
config = context.config
section = config.config_ini_section

config.set_section_option(section, "DB_NAME", settings.db.name)
config.set_section_option(section, "DB_USER", settings.db.user)
config.set_section_option(section, "DB_PASSWORD", settings.db.password)
config.set_section_option(section, "DB_HOST", settings.db.host)
config.set_section_option(section, "DB_PORT", str(settings.db.port))

# Interpret the config file for Python logging.
# This line sets up loggers basically.
if config.config_file_name is not None:
    fileConfig(config.config_file_name)

# add your model's MetaData object here
# for 'autogenerate' support
# from myapp import mymodel
# target_metadata = mymodel.Base.metadata
target_metadata = Base.metadata

# other values from the config, defined by the needs of env.py,
# can be acquired:
# my_important_option = config.get_main_option("my_important_option")
# ... etc.


def run_migrations_offline() -> None:
    """Запуск миграций в режиме 'offline'."""
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
    """Запуск миграций в режиме 'online'."""
    with engine.connect() as connection:
        context.configure(connection=connection, target_metadata=target_metadata)

        with context.begin_transaction():
            context.run_migrations()


if context.is_offline_mode():
    run_migrations_offline()
else:
    run_migrations_online()
