package config

import (
	"github.com/spf13/viper"
)

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./package/config")

	err := viper.ReadInConfig()
	return err
}

func InitForTests() error {
	viper.SetConfigName("config-test")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./package/config")

	err := viper.ReadInConfig()
	return err
}
