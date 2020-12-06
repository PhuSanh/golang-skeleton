package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ApiPort  string   `mapstructure:"API_PORT"`
	Database Database `mapstructure:",squash"`
	Jwt      Jwt      `mapstructure:",squash"`
	Redis    Redis    `mapstructure:",squash"`
}

type Database struct {
	Host string `mapstructure:"DB_HOST"`
	User string `mapstructure:"DB_USERNAME"`
	Pass string `mapstructure:"DB_PASSWORD"`
	Port string `mapstructure:"DB_PORT"`
	Name string `mapstructure:"DB_DATABASE"`
}

type Jwt struct {
	Secret         string `mapstructure:"JWT_SECRET"`
	ExpireInMinute int64  `mapstructure:"JWT_EXPIRE_IN_MINUTE"`
}

type Redis struct {
	Host     string `mapstructure:"REDIS_HOST"`
	Port     string `mapstructure:"REDIS_PORT"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	Pool     int    `mapstructure:"REDIS_POOL"`
}

func LoadConfig() (config *Config, err error) {
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}

	if err = viper.Unmarshal(&config); err != nil {
		return
	}
	return
}
