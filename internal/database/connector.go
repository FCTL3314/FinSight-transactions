package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Connector[T any] interface {
	Connect() (T, error)
}

type pgxConnector struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
}

func NewPgxConnector(
	name, user, password, host, port string,
) Connector[*sql.DB] {
	return &pgxConnector{
		Name: name, User: user, Password: password, Host: host, Port: port,
	}
}

func (c *pgxConnector) Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.Host, c.User, c.Password, c.Name, c.Port)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
