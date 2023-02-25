package util

import "github.com/spf13/viper"

type AppConfig struct {
	AppHost string `mapstructure:"DEV_APP_HOST"`
	AppPort string `mapstructure:"DEV_APP_PORT"`
}
type DBConfig struct {
	Driver   string `mapstructure:"DEV_DB_DRIVER"`
	Host     string `mapstructure:"DEV_DB_HOST"`
	Port     string `mapstructure:"DEV_DB_PORT"`
	User     string `mapstructure:"DEV_DB_USER"`
	Password string `mapstructure:"DEV_DB_PASSWORD"`
	Name     string `mapstructure:"DEV_DB_NAME"`
}

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&config)
	return
}
