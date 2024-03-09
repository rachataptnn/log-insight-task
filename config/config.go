package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	LogFilePath string `mapstructure:"logFilePath"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
