package utils

import (
	"github.com/JeremyLoy/config"
	"github.com/labstack/gommon/log"
)

type EnvConfig struct {
	ServerPort             int    `config:"SERVER_PORT"`
	RepositoryHost         string `config:"REPOSITORY_HOST"`
	RepositoryPort         int    `config:"REPOSITORY_PORT"`
	RepositoryUsername     string `config:"REPOSITORY_USERNAME"`
	RepositoryPassword     string `config:"REPOSITORY_PASSWORD"`
	RepositoryDatabaseName string `config:"REPOSITORY_DATABASE_NAME"`
}

var ec EnvConfig

func GetEnvConfig() *EnvConfig {
	err := config.FromEnv().To(&ec)
	if err != nil {
		log.Fatal(err)
	}
	return &ec

}
