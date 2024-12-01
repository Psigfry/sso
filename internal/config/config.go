package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env      string        `yaml:"env" env-default:"local"`
	Storage  string        `yaml:"storage_path" env-required:"true"`
	TokenTTL time.Duration `yaml:"token_ttl" env-required:"true"`
	GRPC     GRPCConfig    `yaml:"grps"`
}

type GRPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()

	if path == "" {
		panic("config path is required")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not found" + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config file: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var path string
	var defaultPath = "./config/local.yaml"

	flag.StringVar(&path, "config", "", "path to config file")
	flag.Parsed()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	if path == "" {
		return defaultPath
	}

	return path
}
