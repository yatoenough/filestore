package config

import (
	"flag"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env     string `yaml:"env_mode" env:"ENV_MODE" json:"envMode" env-default:"development"`
	ConnStr string `yaml:"pg_connection_string" env:"PG_CONNECTION_STRING" json:"pgConnectionString" env-required:"true"`
	Address string `yaml:"address" env:"ADDRESS" json:"address" env-default:"localhost:3000"`
}

func MustLoad() *Config {
	cfgPath := flag.String("cfgpath", "", "Path to config file. Example: -cfgpath=./configs/cfg.yaml")
	flag.Parse()

	if *cfgPath == "" {
		panic("-cfgpath is not set. Please provide path to a config file with -cfgpath. Example: -cfgpath=./configs/cfg.yaml")
	}

	if _, err := os.Stat(*cfgPath); os.IsNotExist(err) {
		panic(fmt.Sprintf("config file does not exist: %s", *cfgPath))
	}

	var cfg Config

	if err := cleanenv.ReadConfig(*cfgPath, &cfg); err != nil {
		panic(fmt.Sprintf("cannot read config: %s", err))
	}

	return &cfg
}
