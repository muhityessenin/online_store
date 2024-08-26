package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func LoadConfig() (cfg Config, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return cfg, nil
	}
	cfg.Host = os.Getenv("DB_HOST")
	cfg.Port = os.Getenv("DB_PORT")
	cfg.User = os.Getenv("DB_USER")
	cfg.Password = os.Getenv("DB_PASSWORD")
	cfg.Name = os.Getenv("DB_NAME")
	if err = envconfig.Process("", &cfg); err != nil {
		return
	}
	return cfg, nil
}
