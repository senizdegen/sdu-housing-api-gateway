package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/senizdegen/sdu-housing/api-gateway/pkg/logging"
)

type Config struct {
	IsDebug         *bool `yaml:"is_debug"`
	JWT             `yaml:"jwt"`
	Listen          `yaml:"listen"`
	UserService     `yaml:"user_service"`
	PropertyService `yaml:"property_service"`
}

type JWT struct {
	Secret string `yaml:"secret" env-required:"true"`
}

type Listen struct {
	Type   string `yaml:"type" env-default:"port"`
	BindIP string `yaml:"bind_ip" env-default:"localhost"`
	Port   string `yaml:"port" env-default:"8080"`
}

type UserService struct {
	URL string `yaml:"url" env-required:"true"`
}
type PropertyService struct {
	URL string `yaml:"url" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application config")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}

	})

	return instance
}
