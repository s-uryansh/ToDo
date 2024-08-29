package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

// Defining errors
var ErrInvalidDatabaseConfig = errors.New("invalid database config")
var ErrInvalidPort = errors.New("invalid server port")

// Defining Structs
type Config struct {
	DBHost          string `mapstructure:"DB_HOST"`
	DBPort          int    `mapstructure:"DB_PORT"`
	DBUser          string `mapstructure:"DB_USERNAME"`
	DBPassword      string `mapstructure:"DB_PASSWORD"`
	DBName          string `mapstructure:"DB_NAME"`
	RedisAddr       string `mapstructure:"REDIS_ADDR"`
	JwtSecret       string `mapstructure:"JWT_SECRET"`
	Port            int    `mapstructure:"PORT"`
	SSOEnabled      bool
	SSOProvider     string
	SSOClientID     string
	SSOClientSecret string
	SSOCallbackURL  string
}

//mapstructure = key for unmarshalling data from map

func LoadConfig() (*Config, error) {

	//setting-up

	viper.AddConfigPath("../../internal/config/")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	// load config from env
	viper.AutomaticEnv()

	// loading  yaml config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Error loading config file:", err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// Validating config
func (c *Config) ValidateConfig() error {
	if c.DBPort == 0 {
		return ErrInvalidPort
	}
	if c.DBUser == "" || c.DBPassword == "" || c.DBHost == "" || c.DBName == "" {
		return ErrInvalidDatabaseConfig
	}
	return nil
}

// Fetch config
func GetConfig() (*Config, error) {
	cfg, err := LoadConfig()
	if err != nil {
		return nil, err
	}
	err = cfg.ValidateConfig()
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
