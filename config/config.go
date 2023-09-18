package config

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		HTTP `yaml:"http"`
	}

	HTTP struct {
		Port string `yaml:"port"`
	}
)

func NewConfig() (*Config, error) {
	cfg := Config{}

	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
