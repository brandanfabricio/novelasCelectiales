package config

import (
	_ "embed"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

//go:embed env.yaml

var configFile []byte

type Database struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
	Name string `yaml:"name"`
}

type Config struct {
	Port string   `yaml:"port"`
	DB   Database `yaml:"database"`
}

func New() (*Config, error) {
	var config Config
	err := yaml.Unmarshal(configFile, &config)

	if err != nil {
		return nil, err
	}

	// Leer configuración de la base de datos desde variables de entorno si están presentes
	if host := os.Getenv("MYSQL_HOST"); host != "" {
		config.DB.Host = host
	}
	if portStr := os.Getenv("MYSQL_PORT"); portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err == nil {
			config.DB.Port = port
		}
	}
	if user := os.Getenv("MYSQL_USER"); user != "" {
		config.DB.User = user
	}
	if pass := os.Getenv("MYSQL_PASSWORD"); pass != "" {
		config.DB.Pass = pass
	}
	if name := os.Getenv("MYSQL_DB"); name != "" {
		config.DB.Name = name
	}

	

	return &config, nil

}
