package config

import (
	"github.com/jinzhu/configor"
	"time"
)

// Config configuration
type Config struct {
	App struct {
		Port         int
		Mode         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}
	DB struct {
		DbType  string
		ConnStr string
	}
}

// YamlConfig configuration
var YamlConfig = Config{}

func init() {
	configor.New(&configor.Config{Verbose: true}).Load(&YamlConfig, "configs/application.yml")
}
