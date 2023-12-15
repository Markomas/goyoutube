package config

import (
	"github.com/mcuadros/go-defaults"
	"github.com/spf13/viper"
	"youtube/internal/app"
	"youtube/internal/db"
)

type Config struct {
	Name     string     `mapstructure:"name" default:"youtube"`
	App      app.Config `mapstructure:"server"`
	Database db.Config  `mapstructure:"database"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config

	err := viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	defaults.SetDefaults(&config)

	return &config, nil
}
