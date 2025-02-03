package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const CONFIG_PATH_FLAG_KEY = "config_path"
const CONFIG_PATH_ENV_KEY = "CONFIG_PATH"

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"true"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	RPC         RPCConfig     `yaml:"rpc"`
}

type RPCConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout" env-required:"true"`
}

func MustLoad() *Config {
	path := loadConfigPath()
	validateConfigPath(path)

	return createConfig(path)
}

// Priority: flag > env > default
func parseConfigPath() string {
	var path string

	flag.StringVar(&path, CONFIG_PATH_FLAG_KEY, "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv(CONFIG_PATH_ENV_KEY)
	}

	return path
}

func loadConfigPath() string {
	path := parseConfigPath()

	if path == "" {
		panic("config path is empty")
	}

	return path
}

func validateConfigPath(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist")
	}
}

func createConfig(path string) *Config {
	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}
