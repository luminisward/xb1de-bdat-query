package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DbConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
}

type Config struct {
	DbConfig DbConfig
	Addr    string
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		v := viper.New()
		v.SetConfigName("config")                // name of config file (without extension)
		v.AddConfigPath(".")                     // optionally look for config in the working directory
		if err := v.ReadInConfig(); err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		if err := v.Unmarshal(&config); err != nil { // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	}
	return config
}
