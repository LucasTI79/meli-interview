package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Host         string
	Port         string
	TimeoutRead  time.Duration
	TimeoutWrite time.Duration
	TimeoutIdle  time.Duration
}

type Config struct {
	Server ServerConfig
}

func LoadConfig() *Config {
	viper.SetDefault("SERVER_HOST", "0.0.0.0")
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("SERVER_TIMEOUT_READ", 5)
	viper.SetDefault("SERVER_TIMEOUT_WRITE", 5)
	viper.SetDefault("SERVER_TIMEOUT_IDLE", 5)

	viper.AutomaticEnv()

	timeoutRead := viper.GetInt("SERVER_TIMEOUT_READ")
	timeoutWrite := viper.GetInt("SERVER_TIMEOUT_WRITE")
	timeoutIdle := viper.GetInt("SERVER_TIMEOUT_IDLE")

	cfg := &Config{
		Server: ServerConfig{
			Host:         viper.GetString("SERVER_HOST"),
			Port:         viper.GetString("SERVER_PORT"),
			TimeoutRead:  time.Duration(timeoutRead) * time.Second,
			TimeoutWrite: time.Duration(timeoutWrite) * time.Second,
			TimeoutIdle:  time.Duration(timeoutIdle) * time.Second,
		},
	}

	log.Printf("Config loaded: %+v\n", cfg)
	return cfg
}
