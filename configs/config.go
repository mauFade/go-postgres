package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	Api      apiConfig
	Database dbConfig
}

type apiConfig struct {
	Port string
}

type dbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	error := viper.ReadInConfig()

	if error != nil {
		if _, ok := error.(viper.ConfigFileNotFoundError); !ok {
			return error
		}
	}

	cfg = new(config)

	cfg.Api = apiConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.Database = dbConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.database"),
	}

	return nil
}

func GetDatabase() dbConfig {
	return cfg.Database
}

func GetApiPort() string {
	return cfg.Api.Port
}
