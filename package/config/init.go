package config

import (
	"os"

	"github.com/spf13/viper"
)

func Init() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./package/config")

	viper.AutomaticEnv()
	if os.Getenv("POSTGRES_HOST") != "" {
		viper.Set("postgres.postgresDsn", "host="+os.Getenv("POSTGRES_HOST")+" port=5432 user="+os.Getenv("POSTGRES_USER")+" password="+os.Getenv("POSTGRES_PASSWORD")+" dbname="+os.Getenv("POSTGRES_DB")+" sslmode=disable")
	}

	err := viper.ReadInConfig()
	return err
}

func InitForTests() error {
	viper.SetConfigName("config-test")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../../package/config")

	err := viper.ReadInConfig()
	return err
}
