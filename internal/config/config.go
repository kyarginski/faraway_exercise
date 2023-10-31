package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env      string `yaml:"env" env-default:"local"`
	Version  string `yaml:"version" env-default:"unknown"`
	TCPPort  int    `yaml:"tcp_port" env-default:"8088"`
	DataFile string `yaml:"data_file" env-default:"config/server/wisdom.txt"`
}

func MustLoad() *Config {
	configPath := os.Getenv("SERVER_CONFIG_PATH")
	if configPath == "" {
		configPath = "config/server/prod.yaml"
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
