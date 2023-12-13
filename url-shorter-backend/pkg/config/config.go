package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	PORT string `mapstructure:"PORT"`
}

func LoadConfig() (*Config, error) {
	var (
		config Config
		err    error
	)
	viper.AddConfigPath("$APP_HOME")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	if err = viper.ReadInConfig(); err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
