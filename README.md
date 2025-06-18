<div align="center">
  <img width="148" height="148" src="https://github.com/user-attachments/assets/b8b8f3ba-d6da-414e-b5f5-339578b498a8"/>
  <h1>FinSight Transactions</h1>
  <p>Backend server for managing financial transactions. Supports transaction tracking, categorization, and financial analytics.</p>

[![Go](https://img.shields.io/badge/Go-1.23.1-45a3ec?style=flat-square)](https://go.dev/)
[![Gin](https://img.shields.io/badge/Gin-1.10.0-458cec?style=flat-square)](https://github.com/gin-gonic/gin)
[![Gorm](https://img.shields.io/badge/Gorm-1.25.12-38B6FF?style=flat-square)](https://github.com/go-gorm/gorm)
[![cleanenv](https://img.shields.io/badge/cleanenv-1.5.0-276867?style=flat-square)](https://github.com/ilyakaznacheev/cleanenv)
</div>

## 🛠️ Tech Stack

- **Backend**: Go with Gin framework
- **Database**: PostgreSQL with GORM ORM
- **Configuration**: cleanenv for environment management
- **Migrations**: Goose for database schema management
- **Containerization**: Docker & Docker Compose

## 📌 Features

- **Core Functionality**:
  - Transaction management (create, read, update, delete)
  - Recurring transaction support
  - Financial analytics (periodic summaries)
  - Pagination with configurable limits

- **Architecture**:
  - Clean architecture with separated layers (controllers, usecases, repositories)
  - Error handling middleware with domain-specific errors
  - Configurable logging with structured JSON output
  - Database abstraction with GORM ORM

- **Infrastructure**:
  - Docker-based development environment
  - Database migrations with Goose
  - Environment configuration with cleanenv
  - Configurable pagination limits

## ⚙️ Configuration

The application uses environment variables and a `config.yml` file for configuration. Example `.env` file:

```env
DB_NAME=postgres
DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5432
PORT=8080
GIN_MODE=debug
TRUSTED_PROXIES=localhost,127.0.0.1
```

## 📦 API Endpoints

### Transactions
- `GET /api/v1/transactions` - List transactions with pagination
- `GET /api/v1/transactions/:id` - Get specific transaction
- `POST /api/v1/transactions` - Create new transaction
- `PUT /api/v1/transactions/:id` - Update transaction
- `DELETE /api/v1/transactions/:id` - Delete transaction

## 🧪 Testing

The application includes comprehensive testing capabilities:

1. Start test database:
```bash
make up_local_services
```

2. Run migrations:
```bash
make apply_migrations
```

3. Run tests:
```bash
go test ./...
```

## 📃 Notes

* All Docker volumes are stored in the `docker/local/volumes/` directory. If you need to reset your database or any other data, simply delete the corresponding folder.

## ⚒️ Development

1. Download dependencies:
```bash
go mod download
```

2. Install goose for migrations:
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

3. Start Docker services:
```bash
make up_local_services
```

4. Run database migrations:
```bash
make apply_migrations
```

5. Start the application:
```bash
make run
```
