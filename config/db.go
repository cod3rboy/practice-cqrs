package config

type DBConfig struct {
	DBHost     string `env:"PSQL_HOST" envDefault:"localhost"`
	DBPort     int    `env:"PSQL_PORT" envDefault:"5432"`
	DBUser     string `env:"PSQL_USER"`
	DBPassword string `env:"PSQL_PASSWORD"`
	Database   string `env:"PSQL_DB"`
}
