package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIPort      string `default:"8080"`
	PgConnString string `default:"postgres://admin:admin@localhost:5432/comments?sslmode=disable"`
	LogLevel     string `default:"debug"`
}

func InitConfig() (*Config, error) {
	var cnf Config

	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	cnf = Config{
		APIPort:      os.Getenv("SERVER_PORT"),
		PgConnString: os.Getenv("PG_DSN"),
	}

	return &cnf, nil
}
