package config

import (
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	DBConfig
	ServerConfig
	TemporalConfig
}

func NewConfig() Config {
	dbConfig := DBConfig{}
	if err := env.Parse(&dbConfig); err != nil {
		log.Fatal("unable to load database configuration from environment")
	}
	serverConfig := ServerConfig{}
	if err := env.Parse(&serverConfig); err != nil {
		log.Fatal("unable to load server configuration from environment")
	}
	temporalConfig := TemporalConfig{}
	if err := env.Parse(&temporalConfig); err != nil {
		log.Fatalf("unable to load temporal configuration from environment")
	}

	return Config{
		DBConfig:       dbConfig,
		ServerConfig:   serverConfig,
		TemporalConfig: temporalConfig,
	}
}
