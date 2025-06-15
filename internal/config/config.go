package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Pagination struct {
	TransactionLimit int `yaml:"transaction_limit"`
}

type Server struct {
	Mode           string   `env:"GIN_MODE" env-default:"debug"`
	TrustedProxies []string `env:"TRUSTED_PROXIES" env-separator:","`
	Port           string   `env:"PORT" env-default:"8080"`
}

type Database struct {
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
}
type Config struct {
	Server
	Database
	Pagination
}

func Load() (*Config, error) {
	var cfg Config

	baseDir, _ := os.Getwd()

	if err := cleanenv.ReadConfig(fmt.Sprintf("%s/.env", baseDir), &cfg); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadConfig(fmt.Sprintf("%s/config.yml", baseDir), &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
