package config

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

// Config - the config struct for global variables
type Config struct {
	Port       string `mapstructure:"PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
}

// Load pulls the config data from the config file
func Load() (Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return Config{}, fmt.Errorf("error loading config file: %s \n", err)
		}

		// No config file found. Using environment variables
		viper.BindEnv("PORT")
		viper.BindEnv("DB_HOST")
		viper.BindEnv("DB_PORT")
		viper.BindEnv("DB_USER")
		viper.BindEnv("DB_PASSWORD")
		viper.BindEnv("DB_NAME")

		for _, key := range viper.AllKeys() {
			val := viper.Get(key)
			viper.Set(key, val)
		}
	}

	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, fmt.Errorf("unmarshalling config file: %w", err)
	}

	validator := validator.New()
	if err := validator.Struct(config); err != nil {
		return Config{}, fmt.Errorf("invalid config file: %w", err)
	}

	return config, nil
}
