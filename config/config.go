package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	AppPort    string `env:"APP_PORT"`
	JwtSecret  string `env:"JWT_SECRET"`
	DbHost     string `env:"DB_HOST"`
	DbPort     string `env:"DB_PORT"`
	DbUsername string `env:"DB_USERNAME"`
	DbPassword string `env:"DB_PASSWORD"`
	DbName     string `env:"DB_NAME"`
	DbPath     string `env:"DB_PATH"`
}

func MustLoad() *Config {
	var config Config
	if err := cleanenv.ReadEnv(&config); err != nil {
		panic("Cannot read .env file")
	}
	fmt.Println(fmt.Sprintf("APP_PORT: %s", config.AppPort))
	return &config
}
