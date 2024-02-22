package config

import "github.com/spf13/viper"

type Config struct {
	POSTGRES_USER     string `mapstructure:"POSTGRES_USER"`
	POSTGRES_PASSWORD string `mapstructure:"POSTGRES_PASSWORD"`
	POSTGRES_DB       string `mapstructure:"POSTGRES_DB"`

	REDIS_PASSWORD string `mapstructure:"REDIS_PASSWORD"`
	REDIS_BD       int    `mapstructure:"REDIS_DB"`

	NEO4J_USER     string `mapstructure:"NEO4J_USER"`
	NEO4J_PASSWORD string `mapstructure:"NEO4J_PASSWORD"`
}

func InitConfig() (config *Config, err error) {
	viper.AddConfigPath("../.")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(config)
	return
}
