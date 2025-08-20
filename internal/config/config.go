package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
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
	BaseDir    string
	Server     Server
	Database   Database
	Pagination Pagination
}

func Load() (*Config, error) {
	var cfg Config

	baseDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	cfg.BaseDir = baseDir

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadConfig(fmt.Sprintf("%s/config.yml", baseDir), &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
