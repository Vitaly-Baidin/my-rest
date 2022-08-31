package config

import (
	"github.com/Vitaly-Baidin/my-rest/pkg/logging"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	Listen  ListenConfig  `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
}

type ListenConfig struct {
	BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
	Port   string `yaml:"port" env-default:"8080"`
}

type StorageConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
