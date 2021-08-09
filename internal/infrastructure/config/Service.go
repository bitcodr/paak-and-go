package config

import (
	"time"

	"github.com/spf13/viper"
)

type Service struct {
	Host         string
	RestPort     string        `mapstructure:"rest_port"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

func setDefaultServiceConfig(v *viper.Viper) {
	v.SetDefault("server.host", "0.0.0.0")
	v.SetDefault("server.rest_port", "8081")
	v.SetDefault("server.read_timeout", 30*time.Second)
	v.SetDefault("server.write_timeout", 30*time.Second)
}
