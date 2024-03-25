package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructue:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
}

var Cfg Config

func LoadConfig() (err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	viper.AutomaticEnv()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&Cfg)
	return
}
