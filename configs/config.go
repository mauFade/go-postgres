package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	api      apiConfig
	database dbConfig
}

type apiConfig struct {
	port string
}

type dbConfig struct {
	host     string
	port     string
	user     string
	password string
	database string
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

	cfg.api = apiConfig{
		port: viper.GetString("api.port"),
	}

	cfg.database = dbConfig{
		host:     viper.GetString("database.host"),
		port:     viper.GetString("database.port"),
		user:     viper.GetString("database.user"),
		password: viper.GetString("database.password"),
		database: viper.GetString("database.database"),
	}

	return nil
}

func getDatabase() dbConfig {
	return cfg.database
}

func getApiPort() string {
	return cfg.api.port
}
