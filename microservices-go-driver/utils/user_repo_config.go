package utils

import "github.com/spf13/viper"

type Config struct {
	DBDriver     string `mapstructure:"DB_DRIVER"`
	DBName       string `mapstructure:"DB_NAME"`
	DBHost       string `mapstructure:"DB_HOST"`
	DBPort       string `mapstructure:"DB_PORT"`
	DBCollection string `mapstructure:"DB_COLLECTION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
