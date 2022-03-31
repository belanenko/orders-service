package sqlstore

type Config struct {
	DatabaseUrl string `env:"DATABASE_URL"`
}

func NewConfig() *Config {
	return &Config{}
}
