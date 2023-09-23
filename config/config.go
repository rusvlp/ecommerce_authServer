package config

import "github.com/ilyakaznacheev/cleanenv"

type (
	Config struct {
		HTTP     `yaml:"http"`
		Database `yaml:"mysql"`
	}

	HTTP struct {
		Port string `yaml:"port"`
	}

	Database struct {
		URL      string `yaml:"url"`
		username string `yaml:"username"`
		password string `yaml:"password"`
	}
)

func (cfg *Config) String() string {
	return ""
}

func NewConfig() (*Config, error) {
	cfg := Config{}

	err := cleanenv.ReadConfig("./config/config.yml", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
