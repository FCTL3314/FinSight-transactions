package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type DatabaseConfig struct {
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     string `envconfig:"DB_PORT" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
}

func (dc *DatabaseConfig) GetDSN() string {
	return "host=" + dc.Host +
		" port=" + dc.Port +
		" user=" + dc.User +
		" password=" + dc.Password +
		" dbname=" + dc.Name +
		" sslmode=disable"
}

type Server struct {
	ServerAddress string `envconfig:"SERVER_ADDRESS" default:"8080"`
}

type Config struct {
	Debug    bool
	Server   Server
	Database DatabaseConfig
}

func Load() (*Config, error) {
	_ = godotenv.Load(".env")

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
