package config

import (
	"fmt"
	"log"
	"sync"

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

var (
	once sync.Once
	Cfg  *Config
)

func MustLoad() *Config {
	once.Do(func() {
		Cfg = &Config{}
		if err := cleanenv.ReadEnv(Cfg); err != nil {
			log.Fatalf("Cannot read .env file: %s", err)
		}
		fmt.Println(fmt.Sprintf("APP_PORT: %s", Cfg.AppPort))
	})
	return Cfg
}

func (cfg *Config) GetDbUrl() string {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		cfg.DbUsername,
		cfg.DbPassword,
		cfg.DbHost,
		cfg.DbName,
	)
	return dbUrl
}
