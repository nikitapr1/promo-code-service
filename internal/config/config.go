package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env         string `env:"ENV"`
	DatabaseUrl string `env:"DATABASE_URL"`
	GRPCPort    int    `env:"GRPC_PORT"`
}

func MustLoad() *Config {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic("failed to read env " + err.Error())
	}

	return &cfg
}
