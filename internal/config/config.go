package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"go-rest-api/pkg/logging"
	"sync"
)

type Config struct {
	IsDebug  *bool                 `yaml:"is_debug" env-required:"true"`
	Listen   Listen                `yaml:"listen"`
	MongoDB  MongoDBStorageConfig  `yaml:"mongodb"`
	Postgres PostgresStorageConfig `yaml:"postgres"`
}

type Listen struct {
	Type   string `yaml:"type" env-default:"port"`
	BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
	Port   string `yaml:"port" env-default:"8080"`
}

type MongoDBStorageConfig struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Database   string `yaml:"database"`
	Collection string `yaml:"collection"`
	AuthDB     string `yaml:"auth_db"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
}

type PostgresStorageConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
