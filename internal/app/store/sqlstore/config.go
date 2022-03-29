package sqlstore

type Config struct {
	DatabaseURL string `env:"PG_DATABASE_URL"`
}

func NewConfig() *Config {
	return &Config{}
}
