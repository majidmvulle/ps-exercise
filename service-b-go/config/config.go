package config

import (
	"github.com/spf13/viper"
)

var cfg *AppConfig

type AppConfig struct {
	App struct {
		Name  string
		Debug bool
		Env   string
	}

	Gin struct {
		Mode string
		Port uint16
	}

	Database struct {
		DSN    string
		Driver string
	}

	ApiKey string
}

func Config() *AppConfig {
	if cfg == nil {
		loadConfig()
	}

	return cfg
}

func loadConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	_ = viper.ReadInConfig()

	cfg = &AppConfig{}

	// App.
	cfg.App.Name = viper.GetString("APP_NAME")
	cfg.App.Debug = viper.GetBool("APP_DEBUG")
	cfg.App.Env = viper.GetString("APP_ENV")

	// Gin.
	cfg.Gin.Mode = viper.GetString("GIN_MODE")
	cfg.Gin.Port = uint16(viper.GetUint("GIN_PORT"))

	// Database.
	cfg.Database.DSN = viper.GetString("DATABASE_DSN")
	cfg.Database.Driver = viper.GetString("DATABASE_DRIVER")

	// API Key.
	cfg.ApiKey = viper.GetString("API_KEY")
}
