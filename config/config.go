package config

import (
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTPServer *HTTPServer `yaml:"http_server"`
	Database   *Database   `yaml:"database"`
}

type HTTPServer struct {
	Port string `yaml:"port"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

func MustLoad() *Config {
	configPath := "./config/config.yaml"

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic(fmt.Sprintf("Config file does not exist: %s", configPath))
	}

	var config Config
	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		panic(fmt.Sprintf("Cannot read config: %s", configPath))
	}

	return &config
}
