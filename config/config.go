package config

import (
	"errors"
	"github.com/luongdev/openacd/infras/logger"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	App      *AppConfig `mapstructure:"app"`
	Database *DbConfig  `mapstructure:"database"`
}

func LoadConfigPath(path string) (*Config, error) {
	var config *Config
	if path == "" {
		return nil, errors.New("config path is empty")
	}

	viper.SetDefault("database.pool_size", 200)
	viper.SetDefault("database.auth_source", "admin")
	viper.SetDefault("database.connect_timeout", time.Second*5)
	viper.SetDefault("database.timeout", time.Second*30)

	log := logger.Default()
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	err := viper.Unmarshal(&config)
	if err != nil {
		log.Errorf("error to decode, %v", err)
		return nil, err
	}

	return config, nil
}

func LoadConfig() *Config {
	config, _ := LoadConfigPath("config.yml")
	if config == nil {
		config = &Config{
			App:      &AppConfig{},
			Database: viper.Get("database").(*DbConfig),
		}
	}

	return config
}
