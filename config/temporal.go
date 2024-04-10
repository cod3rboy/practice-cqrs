package config

type TemporalConfig struct {
	TemporalHost      string `env:"TEMPORAL_HOST" envDefault:"localhost"`
	TemporalPort      int    `env:"TEMPORAL_PORT" envDefault:"7233"`
	TemporalNamespace string `env:"TEMPORAL_NAMESPACE" envDefault:"default"`
}
