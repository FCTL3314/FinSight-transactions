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
