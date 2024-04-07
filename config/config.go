package config

import (
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	DBConfig
	ServerConfig
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

	return Config{
		DBConfig:     dbConfig,
		ServerConfig: serverConfig,
	}
}
