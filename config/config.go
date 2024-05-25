package config

import (
	"log"

	"github.com/spf13/viper"
)

var Values *Config

type Config struct {
	MySqlUrl string `mapstructure:"MYSQL_URL"`
	HttpPort string `mapstructure:"HTTP_PORT"`
}

func LoadConfig() (config *Config) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("Error reading config file or maybe it was not present: ", err.Error())
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalln("Unable to decode into struct, ", err.Error())
	}

	Values = config
	return config
}
