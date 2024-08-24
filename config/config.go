package config

import (
	"bytes"
	"embed"
	"log"

	"github.com/spf13/viper"
)

//go:embed config.yaml
var configFile embed.FS

type Config struct {
	Database struct {
		DataSourceName string
	}
}

var Cfg *Config

func init() {
	viper.SetConfigType("yaml")

	configContent, err := configFile.ReadFile("config.yaml")
	if err != nil {
		log.Fatalf("Failed to read embedded config file: %v", err)
	}

	if err := viper.ReadConfig(bytes.NewBuffer(configContent)); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	Cfg = &Config{}
	if err := viper.Unmarshal(Cfg); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}
}
