package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connector[T any] interface {
	Connect() (T, error)
}

type gormConnector struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
}

func NewGormConnector(
	name string,
	user string,
	password string,
	host string,
	port string,
) Connector[*gorm.DB] {
	return &gormConnector{
		Name:     name,
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
	}

}

func (c *gormConnector) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		c.Host,
		c.User,
		c.Password,
		c.Name,
		c.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
