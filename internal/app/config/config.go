package config

import (
	"fmt"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"development"`
	ConnStr    string `yaml:"pg_connection_string"  env-required:"true"`
	HTTPServer `yaml:"server_config"`
}

type HTTPServer struct {
	Address        string        `yaml:"address" env-default:"localhost:3000"`
	RequestTimeout time.Duration `yaml:"req_timeout" env-default:"4s"`
	IdleTimeout    time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		panic("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		panic(fmt.Sprintf("config file does not exist: %s", cfgPath))
	}

	var cfg Config

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		panic(fmt.Sprintf("cannot read config: %s", err))
	}

	return &cfg
}
