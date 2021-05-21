package config

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB			*DBConfig
	App			*AppCfg
	RabbitMQ	*RabbitCfg
}

type DBConfig struct {
	URI		string 		`envconfig:"CP_DATABASE_URL"`
}

type AppCfg struct {
	Port	string		`envconfig:"CP_APP_PORT"`
}

type RabbitCfg struct {
	URL		string		`envconfig:"CP_AMQP_SERVER_URL"`
}

func GetConfig() *Config {
	var config Config

	if err := godotenv.Load("./.env"); err != nil {
		log.Panic("Failed to get config")
	}

	if err := envconfig.Process("cp", &config); err != nil {
		log.Panic("Faield to parse config from env", err)
	}

	return &config
}