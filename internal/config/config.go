package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	DbURL string `env:"DATABASE_URL"`
	Port  int    `env:"SERVER_PORT"`
}

func MustLoad() *Config {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("failed to read config: %s", err)
	}
	return &cfg
}
