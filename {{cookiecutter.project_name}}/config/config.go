package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init() *viper.Viper {

	env := os.Getenv("SCOPE")
	if env == "" {
		env = "local"
	}

	config := viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName(env)
	config.AddConfigPath("config/resource/")
	err := config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file", err)
	}
	return config
}
