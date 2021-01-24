package config

import "github.com/spf13/viper"

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../")
	return viper.ReadInConfig()
}
