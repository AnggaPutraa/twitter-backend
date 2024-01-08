package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost             string `mapstructure:"DB_HOST"`
	DBPort             int    `mapstructure:"DB_PORT"`
	DBUsername         string `mapstructure:"DB_USERNAME"`
	DBPassword         string `mapstructure:"DB_PASSWORD"`
	DBDatabase         string `mapstructure:"DB_DATABASE"`
	AccessTokenSecret  string `mapstructure:"ACCESS_TOKEN_SECRET"`
	RefreshTokenSecret string `mapstructure:"REFRESH_TOKEN_SECRET"`
	DBUrl              string
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err = viper.Unmarshal(&config); err != nil {
		return
	}
	config.DBUrl = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBDatabase,
	)
	return config, nil
}
