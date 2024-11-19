package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	MongoDBString string `mapstructure:"mongodb"`
}

// InitConfig initializes the global configuration object
func NewConfig() *Config {
	viper.SetConfigName("configs/local")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err.Error())
	}

	config := Config{}
	viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("unable to marshal config in struct")
	}

	return &config
}
