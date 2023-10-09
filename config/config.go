package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress  string   `mapstructure:"server_address"`
	ContextTimeout int      `mapstructure:"context_timeout"`
	DBConfig       DBConfig `mapstructure:"db_config"`
}

type DBConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
	Name string `mapstructure:"name"`
}

func NewConfig() *Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	var c *Config
	err = viper.Unmarshal(&c)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	return c
}
