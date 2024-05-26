package config

import (
	"github.com/achintya-7/dating-app/logger"
	"github.com/spf13/viper"
)

var Values *Config

type Config struct {
	MySqlUrl          string `mapstructure:"MYSQL_URL"`
	RedisUrl          string `mapstructure:"REDIS_URL"`
	HttpPort          string `mapstructure:"HTTP_PORT"`
	TokenSymmetricKey string `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	EmailName         string `mapstructure:"EMAIL_NAME"`
	EmailAddress      string `mapstructure:"EMAIL_ADDRESS"`
	EmailPassowrd     string `mapstructure:"EMAIL_PASSWORD"`
}

func LoadConfig() (config *Config) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal(nil, "Error reading config file or maybe it was not present: ", err.Error())
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		logger.Fatal(nil, "Unable to decode into struct, ", err.Error())
	}

	logger.Info(nil, "Config loaded successfully")

	Values = config
	return config
}
