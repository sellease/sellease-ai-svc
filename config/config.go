package config

import (
	"log"
	"sync"

	"github.com/go-playground/validator"
	"github.com/spf13/viper"
)

type EnvConfig struct {
	Environment         string `validate:"required" mapstructure:"ENVIRONMENT"`
	ServerPort          string `validate:"required" mapstructure:"SERVER_PORT"`
	PostgresHost        string `validate:"required" mapstructure:"POSTGRES_HOST"`
	PostgresDB          string `validate:"required" mapstructure:"POSTGRES_DB"`
	PostgresSchema      string `validate:"required" mapstructure:"POSTGRES_SCHEMA"`
	PostgresUsername    string `validate:"required" mapstructure:"POSTGRES_USERNAME"`
	PostgresPassword    string `validate:"required" mapstructure:"POSTGRES_PASSWORD"`
	PostgresPort        string `validate:"required" mapstructure:"POSTGRES_PORT"`
	JWTSecret           string `validate:"required" mapstructure:"JWT_SECRET"`
	TextCortexAPIKey    string `validate:"required" mapstructure:"TEXT_CORTEX_API_KEY"`
	TextCortexUrl       string `validate:"required" mapstructure:"TEXT_CORTEX_URL"`
	SellEaseAPISvcUrl   string `validate:"required" mapstructure:"SELLEASE_API_SVC_URL"`
	RapidAPIKey         string `validate:"required" mapstructure:"RAPID_API_KEY"`
	RapidAPIKeywordHost string `validate:"required" mapstructure:"RAPID_API_KEYWORD_HOST"`
	RapidAPIKeywordUrl  string `validate:"required" mapstructure:"RAPID_API_KEYWORD_URL"`
}

var (
	config *EnvConfig
	once   sync.Once
)

func init() {
	once.Do(func() {
		viper.AutomaticEnv()
		viper.SetConfigFile(".env")
		config = new(EnvConfig)
		if err := viper.ReadInConfig(); err != nil {
			log.Printf("error reading config - %s", err)

		}
		if err := viper.Unmarshal(config); err != nil {
			log.Printf("unable to decode config - %v", err)
		}

		validate := validator.New()
		if err := validate.Struct(config); err != nil {
			log.Printf("error in config validation - %v", err)
			panic(err)
		}
	})
}

func GetConfig() *EnvConfig {
	return config
}
