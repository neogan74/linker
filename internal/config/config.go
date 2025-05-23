package config

type Config struct {
	BotServerPort     string
	ScapperServerPort string
}

func NewConfig() *Config {
	return &Config{
		BotServerPort:     "8080",
		ScapperServerPort: "8081",
	}
}
