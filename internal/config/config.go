package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const (
	pathEnv="/home/dunice/Рабочий стол/go_second_microservice/.env"
)

type Config struct {
	PORT string `env:"PORT"`
}

func MustLoad(log *logrus.Logger) (*Config, error) {
	const op = "config.config.mustLoad"
	err := godotenv.Load(pathEnv)
	if err != nil {
		log.Fatalf("error during initialization ENV! %s: %w", op, err.Error())
		return nil, err
	}
	return &Config{
		PORT: getEnv("PORT"),
	}, nil
}

func getEnv(key string) string {
	return os.Getenv(key)
}