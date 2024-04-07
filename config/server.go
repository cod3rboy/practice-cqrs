package config

type ServerConfig struct {
	ServerPort  int    `env:"SERVER_PORT" envDefault:"3001"`
	Environment string `env:"ENV" envDefault:"dev"`
}
