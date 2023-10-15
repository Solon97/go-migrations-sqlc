package configs

import (
	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBUrl    string `mapstructure:"DB_URL"`
	DBName   string `mapstructure:"DB_NAME"`
}

func LoadConfig(path string) (*config, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return cfg, nil
}
