package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/mcuadros/go-defaults"
)

type Config struct {
	MySql     MySqlConfig
	ApiServer ApiServerConfig
}

type MySqlConfig struct {
	Host     string `env:"MSQL_HOST"        default:"localhost"`
	Port     string `env:"MYSQL_PORT"        default:"3306"`
	UserName string `env:"MYSQL_USERNAME"        default:"test"`
	Password string `env:"MYSQL_PASSWORD"        default:"123456"`
	DbName   string `env:"DB_NAME"        default:"banking"`
}

type ApiServerConfig struct {
	Host string `env:"HOST"        default:"localhost"`
	Port string `env:"PORT"        default:"8080"`
}

func New() (*Config, error) {
	cfg := new(Config)

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	defaults.SetDefaults(cfg)

	return cfg, nil
}
