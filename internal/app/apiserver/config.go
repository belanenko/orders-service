package apiserver

type Config struct {
	BindAddr string `env:"BIND_ADDR"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}
